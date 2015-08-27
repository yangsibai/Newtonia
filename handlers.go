package main

import (
	//"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Search(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Search")
}
