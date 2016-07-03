package main

import (
	"golang.org/x/crypto/sha3"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (user User) checkPassword(possiblePassword string) bool {
	if user.Password == possiblePassword {
		return true
	}

	return false
}

func encryptPassword(password string) {
	hash := sha3.New256()
	hash.Write([]byte(password))
}

func main() {
	m := martini.Classic()

	db, err := gorm.Open("postgres", "user=postgres dbname=postgresdb password=superuser1.")
	if err != nil {
		panic("Failed to connect database")
	}

	var user User

	db.LogMode(true)
	db.AutoMigrate(&User{})

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

		db.Where("name = ?", username).Find(&user)

		if user.Name != username {
			return http.StatusNotFound, "User with current 'username' not found"
		}

		user.checkPassword(password)
		encryptPassword(password)

		return http.StatusCreated, "Created"
	})

	m.Run()
}