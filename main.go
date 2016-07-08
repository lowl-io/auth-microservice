package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/go-martini/martini"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/sha3"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bytes"
)

func (user User) checkPassword(possiblePassword string) bool {
	if user.Password == possiblePassword {
		hash := sha3.New256()
		hash.Write([]byte(possiblePassword))
		return true
	}
	return false
}

func jsonResponse(response interface{}, w http.ResponseWriter) {

	json, _ :=  json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func loginHandler(response http.ResponseWriter, request *http.Request) (int, string) {
	var dbConfig DataBaseConfig

	jsonStream, error := ioutil.ReadFile("src/main/configs/database.json")
	if error != nil {
		return -1, "Error reading 'database.json' file"
	}

	json.NewDecoder(bytes.NewReader(jsonStream)).Decode(&dbConfig)

	db, error := gorm.Open(dbConfig.Dialect, dbConfig.DataBaseInfo)
	if error != nil {
		return -1, "Failed to connection database"
	}

	//db.LogMode(true)
	db.AutoMigrate(&User{})

	error = request.ParseForm()
	if error != nil {
		return http.StatusBadRequest, "Incorrect POST request"
	}

	username := request.Form.Get("username")
	password := request.Form.Get("password")

	if username == "" || password == "" {
		return http.StatusBadRequest, "Parameter 'username' or 'password' is not valid"
	}

	var user User

	db.Where("name = ?", username).Find(&user)

	if user.Name != username {
		return http.StatusNotFound, "User with current 'username' not found"
	}

	if !user.checkPassword(password) {
		return http.StatusForbidden, "Parametr 'username' or 'password' is not valid"
	}

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid" : user.ID,
	})

	var key JWTKeyConfig

	jsonStream, error = ioutil.ReadFile("src/main/configs/key.json")
	if error != nil {
		return -1, "Error reading 'key.json' file"
	}

	json.NewDecoder(bytes.NewReader(jsonStream)).Decode(&key)

	tokenString, error := signer.SignedString([]byte(key.Key))
	if error != nil {
		return http.StatusBadRequest, "Error when signing token"
	}

	token := Token{tokenString}
	jsonResponse(token, response)

	return http.StatusCreated, ""
}

func main() {
	m := martini.Classic()

	m.Post("/login", loginHandler)

	m.Run()
}