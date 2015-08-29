// https://developers.google.com/web-search/docs/reference#_intro_fonje

package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
