package main

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"net/http"
)

type User struct {
	id       uint
	name     string `grom:"index"`
	password string
	email    string `gorm:"index"`
}

type UserStatus struct {
	isActive  bool
	isBlocked bool
	isDeleted bool
}

func (user User) checkPassword(possiblePassword string) bool {
	if user.password == possiblePassword {
		return true
	}

	return false
}

func main() {
	m := martini.Classic()

	db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm password=mypassword")
	if err != nil {
		panic("Faild to connect database")
	}

	db.AutoMigrate(&User{})

	user := new(User)

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

		if user.name == username {
			user.checkPassword(password)
		}

		return http.StatusCreated, "Created"
	})

	m.Run()
}