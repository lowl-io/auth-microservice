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

	if login == "" || password == "" {
		t.Fatal("Login or password is not valid")
	}

	if strings.Contains(login, "@") {
		db.Where("email = ?", login).Find(&user)
		if user.Email != login {
			t.Fatal("User with current 'login' not found")
		}
	} else {
		db.Where("username = ?", login).Find(&user)
		if user.Username != login {
			t.Fatal("User with current 'login' not found")
		}
	}

	if (login == user.Username || login == user.Email) && password != user.Password {
		t.Fatal("Login or password is not valid")
	}

}
