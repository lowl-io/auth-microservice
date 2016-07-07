package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/go-martini/martini"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/sha3"
	"github.com/jinzhu/gorm"
	"path/filepath"
	"io/ioutil"
	"net/http"
	"fmt"
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

func loginHandler(response http.ResponseWriter, request *http.Request) (int, string) {
	pathToDataBase, _ := filepath.Abs("src/main/configs/database.txt")
	dbConfig, error := ioutil.ReadFile(pathToDataBase)
	if error != nil {
		return -1, "Error reading 'database.txt' file"
	}

	db, error := gorm.Open("postgres", string(dbConfig))
	if error != nil {
		panic("Failed to connect database")
	}

	db.LogMode(true)
	db.AutoMigrate(&User{})

	error = request.ParseForm()
	if error != nil {
		return http.StatusBadRequest, "Incorrect POST request"
	}

	username := request.Form.Get("username")
	password := request.Form.Get("password")

	if username == "" || password == "" {
		return http.StatusBadRequest, "Parametr 'username' or 'password' is not valid"
	}

	var user User

	db.Where("name = ?", username).Find(&user)

	if user.Name != username {
		return http.StatusNotFound, "User with current 'username' not found"
	}

	if !user.checkPassword(password) {
		return http.StatusForbidden, "Parametr 'username' or 'password' is not valid"
	}

	encryptPassword(password)

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid" : user.ID,
	})

	pathToKey, _ := filepath.Abs("src/main/configs/key.txt")
	secret, error := ioutil.ReadFile(pathToKey)
	if error != nil {
		return -1, "Error reading 'key.txt' file"
	}

	tokenString, error := signer.SignedString([]byte(secret))
	if error != nil {
		return http.StatusBadRequest, "Error when signing token"
	}

	fmt.Println(tokenString)

	return http.StatusCreated, "Created"
}

func main() {
	m := martini.Classic()

	m.Post("/login", loginHandler)

	m.Run()
}