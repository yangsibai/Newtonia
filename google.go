package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type SearchResult struct {
	ResponseData struct {
		Results []struct {
			GsearchResultClass string `json:"GsearchResultClass"`
			UnescapedUrl       string `json:"unescapedUrl"`
			Url                string `json:"url"`
			VisibleUrl         string `json:"visibleUrl"`
			CacheUrl           string `json:"cacheUrl"`
			Title              string `json:"title"`
			TitleNoFormatting  string `json:"titleNoFormatting"`
			Content            string `json:"content"`
		} `json:"results"`
		Cursor struct {
			ResultCount string `json:"resultCount"`
			Pages       []struct {
				Start string `json:"start"`
				Label uint16 `json:"label"`
			} `json:"pages"`
			EstimatedResultCount string `json:"estimatedResultCount"`
			CurrentPageIndex     uint16 `json:"currentPageIndex"`
			MoreResultsUrl       string `json:"moreResultsUrl"`
			SearchResultTime     string `json:"searchResultTime"`
		} `json:"cursor"`
	} `json:"responseData"`
	ResponseStatus uint16 `json:"responseStatus"`
}

func GoogleSearch(word string) (error, SearchResult) {
	var result SearchResult
	switch env := os.Getenv("go_env"); env {
	case "DEVELOPMENT":
		return TestGoogleSearch(word)
	case "PRODUCTION":
		url := "https://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&hl=zh&q=" + word
		res, err := http.Get(url)
		if err != nil {
			return err, result
		}
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close() // must close the body when finished with it
		if err != nil {
			return err, result
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return err, result
		}
		return nil, result
	default:
		return errors.New("invalid go_env set " + env), result
	}
}

func TestGoogleSearch(word string) (error, SearchResult) {
	byt := []byte(`
	{
		"responseData": {
			"results":[
			{
				"GsearchResultClass":"GwebSearch",
				"unescapedUrl":"https://nodejs.org/",
				"url":"https://nodejs.org/",
				"visibleUrl":"nodejs.org",
				"cacheUrl":"http://www.google.com/search?q\u003dcache:P2E8k6j2VqoJ:nodejs.org",
				"title":"\u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e",
				"titleNoFormatting":"Node.js",
				"content":"Event-driven I/O server-side JavaScript environment based on V8. Includes API \ndocumentation, change-log, examples and announcements."
			},
			{
				"GsearchResultClass":"GwebSearch",
				"unescapedUrl":"https://nodejs.org/download/",
				"url":"https://nodejs.org/download/",
				"visibleUrl":"nodejs.org",
				"cacheUrl":"http://www.google.com/search?q\u003dcache:MU5V96u72mQJ:nodejs.org",
				"title":"Download - \u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e",
				"titleNoFormatting":"Download - Node.js",
				"content":"Downloads. Download the \u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e source code or a pre-built installer for your \nplatform, and start developing today. Current version: v0.12.7. Windows Installer\n ..."
			},
			{
				"GsearchResultClass":"GwebSearch",
				"unescapedUrl":"http://nqdeng.github.io/7-days-nodejs/",
				"url":"http://nqdeng.github.io/7-days-nodejs/",
				"visibleUrl":"nqdeng.github.io",
				"cacheUrl":"http://www.google.com/search?q\u003dcache:ylfgZVSmmHIJ:nqdeng.github.io",
				"title":"七天学会\u003cb\u003eNodeJS\u003c/b\u003e",
				"titleNoFormatting":"七天学会NodeJS",
				"content":"豆知识： process 是一个全局变量，可通过 process.argv 获得命令行参数。由于 argv\n[0] 固定等于\u003cb\u003eNodeJS\u003c/b\u003e执行程序的绝对路径， argv[1] 固定等于主模块的绝对路径， ..."
			},
			{
				"GsearchResultClass":"GwebSearch",
				"unescapedUrl":"https://github.com/joyent/node/wiki/installation",
				"url":"https://github.com/joyent/node/wiki/installation",
				"visibleUrl":"github.com",
				"cacheUrl":"http://www.google.com/search?q\u003dcache:X7H2TJXwnn0J:github.com",
				"title":"Installation · joyent/\u003cb\u003enode\u003c/b\u003e Wiki · GitHub",
				"titleNoFormatting":"Installation · joyent/node Wiki · GitHub",
				"content":"Aug 20, 2015 \u003cb\u003e...\u003c/b\u003e You can install a pre-built version of \u003cb\u003enode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e via the downloads page available in \na .tar.gz. Or you can use the automatic bash Installer."
			},
			{
				"GsearchResultClass":"GwebSearch",
				"unescapedUrl":"https://github.com/joyent/node/wiki/installing-node.js-via-package-manager",
				"url":"https://github.com/joyent/node/wiki/installing-node.js-via-package-manager",
				"visibleUrl":"github.com",
				"cacheUrl":"http://www.google.com/search?q\u003dcache:B95AIYYt_JcJ:github.com",
				"title":"Installing \u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e via package manager · joyent/node Wiki · GitHub",
				"titleNoFormatting":"Installing Node.js via package manager · joyent/node Wiki · GitHub",
				"content":"Jul 18, 2015 \u003cb\u003e...\u003c/b\u003e Note: The packages on this page are maintained and supported by their \nrespective packagers, not the \u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e core team. Please report any ..."
			},
			{
				"GsearchResultClass":"GwebSearch",
				"unescapedUrl":"https://zh.wikipedia.org/zh/Node.js",
				"url":"https://zh.wikipedia.org/zh/Node.js",
				"visibleUrl":"zh.wikipedia.org",
				"cacheUrl":"http://www.google.com/search?q\u003dcache:MRfDZ-3hpP0J:zh.wikipedia.org",
				"title":"\u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e - 维基百科，自由的百科全书",
				"titleNoFormatting":"Node.js - 维基百科，自由的百科全书",
				"content":"\u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e 是一个开放源代码、跨平台的、用于服务器端和网络应用的运行环境。\u003cb\u003eNode\u003c/b\u003e.\n\u003cb\u003ejs\u003c/b\u003e应用用JavaScript语言写成，在\u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e运行时运行。它支持OS X、Microsoft ..."
			},
			{
				"GsearchResultClass":"GwebSearch",
				"unescapedUrl":"https://en.wikipedia.org/wiki/Node.js",
				"url":"https://en.wikipedia.org/wiki/Node.js",
				"visibleUrl":"en.wikipedia.org",
				"cacheUrl":"http://www.google.com/search?q\u003dcache:r3FvjQMmVegJ:en.wikipedia.org",
				"title":"\u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e - Wikipedia, the free encyclopedia",
				"titleNoFormatting":"Node.js - Wikipedia, the free encyclopedia",
				"content":"\u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e is an open source, cross-platform runtime environment for server-side \nand networking applications. \u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e applications are written in JavaScript and ..."
			},
			{
				"GsearchResultClass":"GwebSearch",
				"unescapedUrl":"http://azure.microsoft.com/en-us/develop/nodejs/",
				"url":"http://azure.microsoft.com/en-us/develop/nodejs/",
				"visibleUrl":"azure.microsoft.com",
				"cacheUrl":"http://www.google.com/search?q\u003dcache:TJQBE5kNrvsJ:azure.microsoft.com",
				"title":"\u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e Developer Center - Microsoft Azure",
				"titleNoFormatting":"Node.js Developer Center - Microsoft Azure",
				"content":"\u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e Developer Center. Windows Azure SDK for \u003cb\u003eNode\u003c/b\u003e.\u003cb\u003ejs\u003c/b\u003e. Support for \ndeploying Windows Azure Websites, Virtual Machines, and Mobile Services, \nconfiguring ..."}],"cursor":{"resultCount":"5,040,000","pages":[{"start":"0","label":1},{"start":"8","label":2},{"start":"16","label":3},{"start":"24","label":4},{"start":"32","label":5},{"start":"40","label":6},{"start":"48","label":7},{"start":"56","label":8}],"estimatedResultCount":"5040000","currentPageIndex":0,"moreResultsUrl":"http://www.google.com/search?oe\u003dutf8\u0026ie\u003dutf8\u0026source\u003duds\u0026start\u003d0\u0026hl\u003dzh-CN\u0026q\u003dnode.js","searchResultTime":"0.24"
			}
		},
		"responseDetails": null,
		"responseStatus": 200
	}
	`)
	var result SearchResult
	err := json.Unmarshal(byt, &result)
	return err, result
}
