package main

import (
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Post("/token", func(req *http.Request) (int, string) {

		username := req.PostFormValue("username")
		password := req.PostFormValue("password")

		if username == "" && password == "" {
			return 400, "Bad Request"
		}

		return 201, "Created"
	})

	m.Run()
}

//POST /token body: username: "", password: ""
//Не пустой приходил ответ 201 CREATED
//
//А если нет username или password или они пустые
//Возвращать BAD_REQUEST