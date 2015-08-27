// https://developers.google.com/web-search/docs/reference#_intro_fonje

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	err, result := search("node")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

type Result struct {
	responseData struct {
		results []struct {
			GsearchResultClass string
			unescapedUrl       string
			url                string
			visibleUrl         string
			cacheUrl           string
			title              string
			titleNoFormatting  string
			content            string
		}
		cursor struct {
			resultCount string
			pages       []struct {
				start string
				label uint16
			}
			estimatedResultCount string
			currentPageIndex     uint16
			moreResultsUrl       string
			searchResultTime     string
		}
	}
	responseStatus uint8
}

func search(word string) (error, Result) {
	var result Result
	url := "https://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&hl=zh&q=" + word
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return err, result
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close() // must close the body when finished with it
	if err != nil {
		return err, result
	}
	fmt.Println(body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err, result
	}
	return nil, result
}
