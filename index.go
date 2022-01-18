package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tradersclub/TCInterview/entiti"
	"github.com/tradersclub/TCInterview/repositories"
	"github.com/tradersclub/TCInterview/service"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/login", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		u := vars["user"]
		p := vars["password"]

		if u != "" {
			if p != "" {
				if service.ProcessUser(u, p) {
					rw.WriteHeader(200)
					json.NewEncoder(rw).Encode("Ok")
				} else {
					http.Error(rw, "Error", 200)
				}
			} else {
				http.Error(rw, "Error", 200)
			}
		} else {
			http.Error(rw, "Error", 200)
		}

	}).Methods("GET")

	r.HandleFunc("/user", func(rw http.ResponseWriter, r *http.Request) {

		var newUser entiti.User
		json.NewDecoder(r.Body).Decode(&newUser)

		if service.CreateUser(newUser) {
			rw.WriteHeader(201)
			json.NewEncoder(rw).Encode("Ok")
		} else {
			http.Error(rw, "Error", 404)
		}

	}).Methods("POST")

	r.HandleFunc("/user/{id}", func(rw http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		i, _ := strconv.Atoi(vars["id"])

		if i != 0 {
			ok, result := repositories.GetUser(i)
			if ok {
				rw.WriteHeader(200)
				json.NewEncoder(rw).Encode(result)
			} else {
				http.Error(rw, "Error", 200)
			}
		} else {
			http.Error(rw, "Error", 200)
		}

	}).Methods("POST")

	http.ListenAndServe(":8080", r)

}
