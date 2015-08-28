package main

import (
	"html/template"
	"net/http"
	"path"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("tmpl", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type PageData struct {
	Title        string
	SearchResult SearchResult
}

func Search(w http.ResponseWriter, r *http.Request) {
	words := r.URL.Query()["q"]
	if words[0] == "" {
		http.Redirect(w, r, "/", 307)
		return
	}
	err, searchResult := GoogleSearch(words[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
