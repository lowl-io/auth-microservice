package main

import (
	"github.com/jinzhu/gorm"
	"testing"
	"strings"
	"fmt"
)

func TestHandler(t *testing.T) {

	db, error := gorm.Open("postgres", "user=postgres dbname=postgresdb sslmode=disable password=")
	if error != nil {
		fmt.Errorf("Failed to connection database")
	}

	user := User{ID: 0, Name: "Alexandr", Username: "s0lus", Email: "example@gmail.com", Password: "password", Status: "active"}

	db.CreateTable(&user)
	db.Create(&user)

	login := "s0lus"
	password := "password"

	// login or password not exists
	if login == "" || password == "" {
		t.Fatal("Login or password is not valid")
	}

	// login = exists, but incorrect
	// password = exists, but incorrect
	if !(login == "") && !(password == "") {
		if strings.Contains(login, "@") {
			db.Where("email = ?", login).Find(&user)
			if login != user.Email {
				t.Fatal("Login or password is not valid by email")
			}
		} else {
			db.Where("username = ?", login).Find(&user)
			if login != user.Username {
				t.Fatal("Login is not valid by username")
			}
		}

		if password != user.Password {
			t.Fatal("Password is not valid")
		}
	}
}