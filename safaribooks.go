package main

import (
	"encoding/json"
	"fmt"
	"github.com/rahulvramesh/oreilly-books-grabber/models"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"strconv"
)

const credentialStore = ".cred"

func main() {

	// login first
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("No commands provided, available commands are")
		fmt.Println("")
		fmt.Println("\tsafaribooks login")
		fmt.Println("\tsafaribooks logout")
		fmt.Println("\tsafaribooks grab <bookId>\n")
		return
	}

	command := args[0]

	switch command {

	case "login":
		_ = os.Remove(credentialStore)
		login()
		break

	case "logout":
		_ = os.Remove(credentialStore)
		fmt.Println("You have successfully logged out.")
		break

	case "grab":

		bookid, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Wrong book id provided.")
			return
		}

		grab(bookid)
		break

	default:
		fmt.Println("Wrong Command, available commands are")
		fmt.Println("")
		fmt.Println("\tsafaribooks login")
		fmt.Println("\tsafaribooks logout")
		fmt.Println("\tsafaribooks grab <bookId>\n")
		os.Exit(0)
		return

	}

}

func login() {

	fmt.Print("Enter the your email address: ")
	var email, password string
	_, err := fmt.Scanf("%s", &email)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Enter the your password: ")
	pass, errpass := terminal.ReadPassword(0)

	password = string(pass)
	fmt.Println()
	if errpass != nil {
		fmt.Println(errpass)
		return
	}

	credentials := models.LoginPayload{email, password}
	credentialjson, _ := json.Marshal(credentials)
	err = ioutil.WriteFile(credentialStore, credentialjson, 0644)

	login, err := getCredential()

	if err != nil {
		fmt.Println("invalid credentials you need to login again.")
		return
	}

	loginRes := models.DoLogin(login)

	if loginRes.IDToken == "" {
		_ = os.Remove("cred.json")
		fmt.Println("login failed, login using \n\n\t safaribooks login \n")
		return
	}

	fmt.Println("\nSuccessfully authenticated.\n")
	fmt.Println("\tsafaribooks grab <bookId> to grab the books\n")

}

func getCredential() (models.LoginPayload, error) {

	plan, _ := ioutil.ReadFile(credentialStore)
	var data models.LoginPayload
	err := json.Unmarshal(plan, &data)
	return data, err
}

func grab(bookid int) {

	login, err := getCredential()

	if err != nil {
		fmt.Println("invalid credentials you need to login again.")
		return
	}

	loginRes := models.DoLogin(login)

	if loginRes.IDToken == "" {
		_ = os.Remove("cred.json")
		fmt.Println("login failed, login using \n\n\t safaribooks login \n")
		return
	}

	//now get the index
	//pass login and book id example: 9781788993708
	bookIndex, _ := models.GetBookIndex(loginRes, bookid)

	//now for each item get content
	models.SaveContentToFile(loginRes, bookIndex)

}
