package main

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"net/http"
)

func (user User) checkPassword(possiblePassword string) bool {
	if user.password == possiblePassword {
		return true
	}

	return false
}

func main() {
	m := martini.Classic()

	db, err := gorm.Open("postgres", "user=postgres dbname=postgresdb password=superuser1.")
	if err != nil {
		panic("Failed to connect database")
	}

	var user User

	db.Where("name = ?", "Alexander").Find(&user)

	m.Post("/token", func(request *http.Request) (int, string) {
		err := request.ParseForm()
		if err != nil {
			return http.StatusBadRequest, "Incorrect POST request"
		}

		username := request.Form.Get("username")
		password := request.Form.Get("password")

		if username == "" || password == "" {
			return http.StatusBadRequest, "Parameter 'username' or 'password' is not valid"
		}

		if user.name == username {
			user.checkPassword(password)
		}

		return http.StatusCreated, "Created"
	})

	m.Run()
}