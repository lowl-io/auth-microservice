package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/go-martini/martini"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/sha3"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"fmt"
	"strings"
)

func (user User) checkPassword(possiblePassword string) bool {
	hash := sha3.New256()
	hash.Write([]byte(possiblePassword))
	possiblePassword = hex.EncodeToString(hash.Sum(nil))

	return user.Password == possiblePassword
}

func jsonResponse(response interface{}, w http.ResponseWriter)  {

	json, error := json.Marshal(response)
	if error != nil {
		fmt.Errorf("Error creating json")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func tokenHandler(response http.ResponseWriter, request *http.Request, config Config, db *gorm.DB) (int, string) {
	error := request.ParseForm()
	if error != nil {
		return http.StatusBadRequest, "Incorrect POST request"
	}

	login := request.Form.Get("username")
	password := request.Form.Get("password")

	if login == "" || password == "" {
		return http.StatusBadRequest, "Parameter 'username' or 'password' is not valid"
	}

	var user User

	if strings.Contains(login, "@") {
		db.Where("email = ?", login).Find(&user)

		if user.Email != login {
			return http.StatusNotFound, "User with current 'username' not found by login"
		}
	} else {
		db.Where("username = ?", login).Find(&user)

		if user.Username != login {
			return http.StatusNotFound, "User with current 'username' not found by login"
		}
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

	//token := Token{tokenString}
	//jsonResponse(token, response)

	return http.StatusCreated, "Token: " + tokenString
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