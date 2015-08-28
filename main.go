// https://developers.google.com/web-search/docs/reference#_intro_fonje

package main

import (
	"fmt"
	"log"
	//"net/http"
)

func main() {
	/*
	 *err, result := search("node")
	 *if err != nil {
	 *log.Fatal(err)
	 *fmt.Println(err)
	 *return
	 *}
	 *fmt.Println(result)
	 */
	err, result := GoogleSearch("node")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	/*
	 *router := NewRouter()
	 *log.Fatal(http.ListenAndServe(":8080", router))
	 *err, result := search("go")
	 *if err != nil {
	 *    log.Fatal(err)
	 *}
	 *fmt.Println(result.responseData)
	 */
}
