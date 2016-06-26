package main

import (
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Post("/token", func(request *http.Request) (int, string) {
		err := request.ParseForm()
		if err != nil {
			return http.StatusBadRequest, "Incorrect POST request"
		}

		username := request.FormValue("username")
		if username == "" {
			return http.StatusBadRequest, "Parameter 'username' is not valid"
		}

		password := request.FormValue("password")
		if password == "" {
			return http.StatusBadRequest, "Parameter 'password' is not valid"
		}

		return http.StatusCreated, "Created"
	})

	m.Run()
}
