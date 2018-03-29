package main

import (
	"encoding/json"
	"html/template"
	"log"
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

type Cursor struct {
	Previous PageInfo
	Current  PageInfo
	Next     PageInfo
}

type PageData struct {
	Title        string
	Keyword      string
	SearchResult SearchResult
	CurrentPage  int
	Cursor       Cursor
}

func Search(w http.ResponseWriter, r *http.Request) {
	word := getQuery(r)
	log.Println(word)
	if word == "" {
		http.Redirect(w, r, "/", 307)
		return
	}
	userip := getUserIP(r)
	startAt := getStart(r)
	lang := getLanguage(r)
	err, searchResult := GoogleSearch(word, startAt, userip, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("tmpl", "search.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var currentPageInfo PageInfo
	var nextPageInfo PageInfo
	var previousPageInfo PageInfo

	if len(searchResult.Queries.Request) > 0 {
		currentPageInfo = searchResult.Queries.Request[0]
	}

	if len(searchResult.Queries.NextPage) > 0 {
		nextPageInfo = searchResult.Queries.NextPage[0]
	}

	if len(searchResult.Queries.PreviousPage) > 0 {
		previousPageInfo = searchResult.Queries.PreviousPage[0]
	}

	cursor := Cursor{previousPageInfo, currentPageInfo, nextPageInfo}

	if currentPageInfo.Count <= 0 {
		b, err := json.Marshal(searchResult)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
		return

	}
	pageData := PageData{word, word, searchResult, currentPageInfo.StartIndex/currentPageInfo.Count + 1, cursor}

	if err := tmpl.Execute(w, pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getUserIP(r *http.Request) string {
	if len(r.Header["X-Forwarded-For"]) != 0 {
		return r.Header["X-Forwarded-For"][0]
	} else {
		return r.RemoteAddr
	}
}

func getQuery(r *http.Request) string {
	words := r.URL.Query()["q"]
	return words[0]
}

func getStart(r *http.Request) int64 {
	start := r.URL.Query()["start"]
	if start != nil && start[0] != "" {
		startAt, _ := strconv.ParseInt(start[0], 10, 32)
		return startAt
	}
	return 1
}

func getLanguage(r *http.Request) string {
	lang := "us"
	languages := r.Header["Accept-Language"]
	if len(languages) > 0 {
		return languages[0]
	}
	return lang
}
