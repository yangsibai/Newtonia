package main

import (
	//"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
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
	CurrentPage  uint16
}

func Search(w http.ResponseWriter, r *http.Request) {
	words := r.URL.Query()["q"]
	if words[0] == "" {
		http.Redirect(w, r, "/", 307)
		return
	}
	var startAt int64 = 0
	var err error
	startAt = 0
	start := r.URL.Query()["start"]
	if start != nil && start[0] != "" {
		if startAt, err = strconv.ParseInt(start[0], 10, 32); err != nil {
			startAt = 0
		}
	}
	err, searchResult := GoogleSearch(words[0], startAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("tmpl", "search.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageData := PageData{words[0], searchResult, searchResult.ResponseData.Cursor.CurrentPageIndex}

	if err := tmpl.Execute(w, pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
