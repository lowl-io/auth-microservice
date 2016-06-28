package main

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

type User struct {
	id       uint
	name     string
	password string
	email    string
}

type UserStatus struct {
	isActive  bool
	isBlocked bool
	isDeleted bool
}

func main() {
	m := martini.Classic()
	
	m.Post("/token", func(request *http.Request) (int, string) {
		err := request.ParseForm()
		if err != nil {
			return http.StatusBadRequest, "Incorrect POST request"
		}

		username := request.Form.Get("username")
		if username == "" {
			return http.StatusBadRequest, "Parameter 'username' is not valid"
		}

		password := request.Form.Get("password")
		if password == "" {
			return http.StatusBadRequest, "Parameter 'password' is not valid"
		}

		return http.StatusCreated, "Created"
	})

	m.Run()
}