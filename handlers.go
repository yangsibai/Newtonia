package main

import (
	//"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

type PageData struct {
	Title        string
	SearchResult SearchResult
}

func Search(w http.ResponseWriter, r *http.Request) {
	words := r.URL.Query()["q"]
	err, searchResult := GoogleSearch(words[0])
	if err != nil {
		fmt.Fprintln(w, "error")
		return
	}

	tmpl, err := template.ParseFiles(path.Join("tmpl", "search.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageData := PageData{words[0], searchResult}
	if err := tmpl.Execute(w, pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
