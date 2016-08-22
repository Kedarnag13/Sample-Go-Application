package main

import (
	"github.com/Kedarnag13/Sample-Go-Application/api/v1/controllers/users"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sign_in", users.Foo.Create).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:3000", nil)
}

// How to Marshal and Unmarshal - https://play.golang.org/p/WxRgpycMaH
