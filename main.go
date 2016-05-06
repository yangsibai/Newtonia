// https://developers.google.com/web-search/docs/reference#_intro_fonje

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Configration struct {
	Port string `json:"port"`
	ID   string `json:"ID"`
	Key  string `json:"key"`
}

var config Configration

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("read config.json error", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("parse config error", err)
	}
}
