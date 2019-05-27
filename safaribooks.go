package main

import (
	"github.com/sivsivsree/oreilly-books-grabber/models"
)

func main() {

	//do login

	login := models.LoginPayload{
		Email:    "email@example.com",
		Password: "password",
	}

	loginRes := models.DoLogin(login)

	if loginRes.IDToken == "" {
		panic("error")
	}

	//now get the index
	//pass login and book id
	bookIndex, _ := models.GetBookIndex(loginRes, 9781788993708)

	//now for each item get content
	models.SaveContentToFile(loginRes, bookIndex)
}
