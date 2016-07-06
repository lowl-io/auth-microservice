package main

import (
	"golang.org/x/crypto/sha3"
	"github.com/go-martini/martini"
	//"github.com/jinzhu/gorm"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"io/ioutil"
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

type token struct {
	Token string `json:"token"`
}

func jsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func loginHandler(response http.ResponseWriter, request *http.Request) {
	var privateKey []byte
	privateKey, _ = ioutil.ReadFile("/home/alex/Workspace/golang/src/main/keys/key.rsa")

	var user User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusForbidden)
		fmt.Println(response, "Error in request")
		return
	}

	fmt.Println(user.ID, user.Name, user.Password)

	if user.ID != 1 {
		response.WriteHeader(http.StatusForbidden)
		fmt.Println("Error logging in")
		fmt.Println(response, "Invalid user info")
		return
	}

	signer := jwt.New(jwt.SigningMethodRS256)

	tokenString, err := signer.SignedString(privateKey)
	if err != nil {
		log.Printf("Error signing token: %v\n", err)
	}

	resp := token{tokenString}
	jsonResponse(resp, response)
}

func main() {
	m := martini.Classic()

	//db, err := gorm.Open("postgres", "user=postgres dbname=postgresdb password=superuser1.")
	//if err != nil {
	//	panic("Failed to connect database")
	//}
	//
	////var user User
	////
	////db.LogMode(true)
	////db.AutoMigrate(&User{})

	m.Post("/login", loginHandler)

	m.Run()
}