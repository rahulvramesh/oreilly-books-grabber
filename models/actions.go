package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//LoginURL -
const LoginURL = "https://www.oreilly.com/member/auth/login/"

//DoLogin
func DoLogin(login LoginPayload) *LoginResponse {

	var loginResponse LoginResponse
	var err error

	b, err := json.Marshal(login)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return &loginResponse
	}

	payload := strings.NewReader(string(b))

	req, _ := http.NewRequest("POST", LoginURL, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "www.oreilly.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &loginResponse)

	if err != nil {
		fmt.Println(err)
		return &loginResponse
	}

	return &loginResponse
}

//GetBookIndex -
func GetBookIndex(login *LoginResponse, bookID int) (BookContext, error) {

	var cover BookContext

	url := fmt.Sprintf("https://learning.oreilly.com/nest/epub/toc/?book_id=%d", bookID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cookie", "orm-jwt="+login.IDToken)
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("authority", "learning.oreilly.com")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "learning.oreilly.com")
	req.Header.Add("Connection", "keep-alive")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//save to the stuct
	err := json.Unmarshal(body, &cover)

	if err != nil {
		fmt.Println(err)
		return cover, err
	}

	return cover, nil

}

//SaveContentToFile -
func SaveContentToFile(login *LoginResponse, bookCover BookContext) {

	for _, bookIndexItem := range bookCover.Items {

		var chapter BookChapter

		url := "https://learning.oreilly.com" + bookIndexItem.URL

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("cookie", "orm-jwt="+login.IDToken)
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
		req.Header.Add("content-type", "application/json; charset=utf-8")
		req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Add("referer", url)
		req.Header.Add("authority", "learning.oreilly.com")
		req.Header.Add("x-requested-with", "XMLHttpRequest")
		req.Header.Add("Authorization", "Bearer "+login.IDToken)
		req.Header.Add("Host", "learning.oreilly.com")
		req.Header.Add("Connection", "keep-alive")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		err := json.Unmarshal(body, &chapter)

		if err != nil {
			fmt.Println("errir")
			fmt.Println(err)
			return
		}

		fmt.Println(chapter.Content)

		getChapterContent(login, chapter, bookIndexItem.Label)

		time.Sleep(5 * time.Second)

	}

}

func getChapterContent(login *LoginResponse, chapter BookChapter, bookCover string) {

	req, _ := http.NewRequest("GET", chapter.Content, nil)

	req.Header.Add("cookie", "orm-jwt="+login.IDToken)
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("authority", "learning.oreilly.com")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("Authorization", "Bearer "+login.IDToken)
	req.Header.Add("Host", "learning.oreilly.com")
	req.Header.Add("Connection", "keep-alive")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	//makefolder
	rootPath := chapter.NaturalKey[0]
	CreateDirIfNotExist(rootPath)

	//save each file
	err := ioutil.WriteFile(rootPath+"/"+bookCover+".html", body, 0644)
	if err != nil {
		panic(err)
	}

	//##mimetype

	createFile(rootPath+"/mimetype", "application/epub+zip")

	//##
	CreateDirIfNotExist(rootPath + "/META-INF")

	optionsXML := `<?xml version="1.0" encoding="UTF-8"?>
		<display_options>
		<platform name="*">
			<option name="specified-fonts">true</option>
		</platform>
	</display_options>
	`

	contentXML := `<?xml version="1.0" encoding="UTF-8"?><container xmlns="urn:oasis:names:tc:opendocument:xmlns:container" version="1.0">
		<rootfiles>
			<rootfile full-path="OEBPS/package.opf" media-type="application/oebps-package+xml"/>
		</rootfiles>
	</container>
	`

	//##
	createFile(rootPath+"/META-INF/com.apple.ibooks.display-options.xml", optionsXML)

	//##

	createFile(rootPath+"/META-INF/container.xml", contentXML)
}

func createFile(path, content string) {

	if _, errF := os.Stat(path); os.IsNotExist(errF) {

		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		w := bufio.NewWriter(f)
		n4, err := w.WriteString(content)
		fmt.Printf("wrote %d bytes\n", n4)

		w.Flush()

	}

}

//CreateDirIfNotExist -
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
