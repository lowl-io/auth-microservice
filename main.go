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
	"fmt"
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

	json, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func tokenHandler(response http.ResponseWriter, request *http.Request, config Config, db *gorm.DB) (int, string) {
	error := request.ParseForm()
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
		return http.StatusForbidden, "Parameter 'username' or 'password' is not valid"
	}

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid" : user.ID,
	})

	tokenString, error := signer.SignedString([]byte(config.JWT.Key))
	if error != nil {
		return http.StatusBadRequest, "Error when signing token"
	}

	token := Token{tokenString}
	jsonResponse(token, response)

	return http.StatusCreated, ""
}

func main() {
	var config Config

	jsonStream, error := ioutil.ReadFile("config.json")
	if error != nil {
		fmt.Errorf("Error reading 'config.json' file")
	}

	json.Unmarshal(jsonStream, &config)

	db, error := gorm.Open(config.DataBase.Dialect, config.DataBase.ConnectionData)
	if error != nil {
		fmt.Errorf("Failed to connection database")
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)

	//db.LogMode(true)
	db.AutoMigrate(&User{})

	m := martini.Classic()

	m.Map(db)
	m.Map(config)

	m.Post("/token", tokenHandler)

	m.Run()
}