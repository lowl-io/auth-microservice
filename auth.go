package main

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

type User struct {
	id       uint
	name     string `grom:"index"`
	password string
	email    string	`gorm:"index"`
}

type UserStatus struct {
	isActive  bool
	isBlocked bool
	isDeleted bool
}

func main() {
	m := martini.Classic()

	db, err := gorm.Open("postgres", "user=postgres dbname=PostgresDB")
	if err != nil {
		panic("Faild to connect database")
	}

	db.AutoMigrate(&User{})

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