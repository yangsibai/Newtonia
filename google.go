package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const apiKey string = "AIzaSyCuCJRJfWuxpNBFi2JJ-c5-lbQ4LZ-eCYk"
const apiID string = "008502474502499729233:r8wzxwoqj9y"

type PageInfo struct {
	Title          string `json:"title"`
	TotalResults   string `json:"totalResults"`
	SearchTerms    string `json:"searchTerms"`
	Count          int    `json:"count"`
	StartIndex     int    `json:"startIndex"`
	InputEncoding  string `json:"inputEncoding"`
	OutputEncoding string `json:"outputEncoding"`
	Safe           string `json:"safe"`
	Cx             string `json:"cs"`
}

type Pagemap struct {
	Metatags []struct {
		Author   string `json:"author"`
		Viewport string `json:"viewport"`
		Date     string `json:"date"`
	} `json:"metatags"`
	Breadcrumb []struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	} `json:"Breadcrumb"`
	Cse_iamge []struct {
		Src string `json:"src"`
	} `json:"cse_image"`
	Cse_thumbnail []struct {
		Width  string `json:"width"`
		Height string `json:"height"`
		Src    string `json:"src"`
	} `json:"cse_thumbnail"`
	Softwaresourcecode []struct {
		Author       string `json:"author"`
		Name         string `json:"name"`
		About        string `json:"about"`
		Url          string `json:"url"`
		Keywords     string `json:"keywords"`
		Datemodified string `json:"datemodified"`
		License      string `json:"license"`
		Text         string `json:"text"`
	} `json:"softwaresourcecode"`
}

type SearchItem struct {
	Kind             string        `json:"kind"`
	Title            string        `json:"title"`
	HtmlTitle        template.HTML `json:"htmlTitle"`
	Link             string        `json:"link"`
	DisplayLink      string        `json:"displayLink"`
	Snippet          string        `json:"snippet"`
	HtmlSnippet      template.HTML `json:"htmlSnippet"`
	CacheId          string        `json:"cacheId"`
	FormattedUrl     string        `json:"formattedUrl"`
	HtmlFormattedUrl template.HTML `json:"htmlFOrmattedUrl"`
	Pagemap          Pagemap       `json:"pagemap"`
}

type SearchResult struct {
	Kind string `json:"kind"`
	url  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	}
	Queries struct {
		Request      []PageInfo `json:"request"`
		NextPage     []PageInfo `json:"nextPage"`
		PreviousPage []PageInfo `json:"previousPage"`
	} `json:"queries"`
	Context struct {
		Title string `json:"title"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float32 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []SearchItem `json:"items"`
}

func GoogleSearch(word string, start int64, userip string, language string) (error, SearchResult) {
	var result SearchResult
	switch env := os.Getenv("go_env"); env {
	case "DEVELOPMENT":
		return TestGoogleSearch(word)
	case "PRODUCTION":
		var Url *url.URL
		Url, err := url.Parse("https://www.googleapis.com/customsearch/v1")
		if err != nil {
			return err, result
		}
		parameters := url.Values{}
		parameters.Add("key", apiKey)
		parameters.Add("cx", apiID)
		parameters.Add("q", word)
		parameters.Add("start", strconv.Itoa(int(start)))
		parameters.Add("hl", language)
		parameters.Add("num", "10")
		parameters.Add("userIp", userip)
		Url.RawQuery = parameters.Encode()
		res, err := http.Get(Url.String())
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
		 "kind": "customsearch#search",
		 "url": {
		  "type": "application/json",
		  "template": "https://www.googleapis.com/customsearch/v1?q={searchTerms}&num={count?}&start={startIndex?}&lr={language?}&safe={safe?}&cx={cx?}&cref={cref?}&sort={sort?}&filter={filter?}&gl={gl?}&cr={cr?}&googlehost={googleHost?}&c2coff={disableCnTwTranslation?}&hq={hq?}&hl={hl?}&siteSearch={siteSearch?}&siteSearchFilter={siteSearchFilter?}&exactTerms={exactTerms?}&excludeTerms={excludeTerms?}&linkSite={linkSite?}&orTerms={orTerms?}&relatedSite={relatedSite?}&dateRestrict={dateRestrict?}&lowRange={lowRange?}&highRange={highRange?}&searchType={searchType}&fileType={fileType?}&rights={rights?}&imgSize={imgSize?}&imgType={imgType?}&imgColorType={imgColorType?}&imgDominantColor={imgDominantColor?}&alt=json"
		 },
		 "queries": {
		  "request": [
		   {
			"title": "Google Custom Search - node",
			"totalResults": "25800000",
			"searchTerms": "node",
			"count": 10,
			"startIndex": 1,
			"inputEncoding": "utf8",
			"outputEncoding": "utf8",
			"safe": "off",
			"cx": "008502474502499729233:r8wzxwoqj9y"
		   }
		  ],
		  "nextPage": [
		   {
			"title": "Google Custom Search - node",
			"totalResults": "25800000",
			"searchTerms": "node",
			"count": 10,
			"startIndex": 11,
			"inputEncoding": "utf8",
			"outputEncoding": "utf8",
			"safe": "off",
			"cx": "008502474502499729233:r8wzxwoqj9y"
		   }
		  ]
		 },
		 "context": {
		  "title": "Newtonia"
		 },
		 "searchInformation": {
		  "searchTime": 0.423118,
		  "formattedSearchTime": "0.42",
		  "totalResults": "25800000",
		  "formattedTotalResults": "25,800,000"
		 },
		 "items": [
		  {
		   "kind": "customsearch#result",
		   "title": "Node.js",
		   "htmlTitle": "\u003cb\u003eNode\u003c/b\u003e.js",
		   "link": "https://nodejs.org/",
		   "displayLink": "nodejs.org",
		   "snippet": "Event-driven I/O server-side JavaScript environment based on V8. Includes API \ndocumentation, change-log, examples and announcements.",
		   "htmlSnippet": "Event-driven I/O server-side JavaScript environment based on V8. Includes API \u003cbr\u003e\ndocumentation, change-log, examples and announcements.",
		   "cacheId": "P2E8k6j2VqoJ",
		   "formattedUrl": "https://nodejs.org/",
		   "htmlFormattedUrl": "https://\u003cb\u003enode\u003c/b\u003ejs.org/",
		   "pagemap": {
			"metatags": [
			 {
			  "author": "Node.js Foundation",
			  "viewport": "width=device-width, initial-scale=1.0"
			 }
			]
		   }
		  },
		  {
		   "kind": "customsearch#result",
		   "title": "Node - YouTube",
		   "htmlTitle": "\u003cb\u003eNode\u003c/b\u003e - YouTube",
		   "link": "https://www.youtube.com/user/Node",
		   "displayLink": "www.youtube.com",
		   "snippet": "Sam, Niko, Brandon and D bringing you quality gaming videos as well as live \naction content every single week! For business inquires please contact us at, inf...",
		   "htmlSnippet": "Sam, Niko, Brandon and D bringing you quality gaming videos as well as live \u003cbr\u003e\naction content every single week! For business inquires please contact us at, inf...",
		   "cacheId": "nsa5oXtjdwwJ",
		   "formattedUrl": "https://www.youtube.com/user/Node",
		   "htmlFormattedUrl": "https://www.youtube.com/user/\u003cb\u003eNode\u003c/b\u003e",
		   "pagemap": {
			"cse_thumbnail": [
			 {
			  "width": "225",
			  "height": "225",
			  "src": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSVdohA6v6X2jZ1JpR1iLiyCuxzbysOcTqzfCW8je32MvRgFMfTDm7Tvvw"
			 }
			],
			"youtubechannelv2": [
			 {
			  "url": "https://www.youtube.com/user/Node",
			  "name": "Node",
			  "description": "Sam, Niko, Brandon and D bringing you quality gaming videos as well as live action content every single week! For business inquires please contact us at, inf...",
			  "paid": "False",
			  "channelid": "UCI4Wh0EQPjGx2jJLjmTsFBQ",
			  "thumbnailurl": "https://yt3.ggpht.com/-217Xf1nmtEc/AAAAAAAAAAI/AAAAAAAAAAA/WCrGMukD8eY/s900-c-k-no-rj-c0xffffff/photo.jpg",
			  "isfamilyfriendly": "True",
			  "regionsallowed": "AD,AE,AF,AG,AI,AL,AM,AO,AQ,AR,AS,AT,AU,AW,AX,AZ,BA,BB,BD,BE,BF,BG,BH,BI,BJ,BL,BM,BN,BO,BQ,BR,BS,BT,BV,BW,BY,BZ,CA,CC,CD,CF,CG,CH,CI,CK,CL,CM,CN,CO,CR,CU,CV,CW,CX,CY,CZ,DE,DJ,DK,DM,DO,DZ,EC,EE,EG,EH..."
			 }
			],
			"imageobject": [
			 {
			  "url": "https://yt3.ggpht.com/-217Xf1nmtEc/AAAAAAAAAAI/AAAAAAAAAAA/WCrGMukD8eY/s900-c-k-no-rj-c0xffffff/photo.jpg",
			  "width": "900",
			  "height": "900"
			 },
			 {
			  "url": "https://yt3.ggpht.com/-217Xf1nmtEc/AAAAAAAAAAI/AAAAAAAAAAA/WCrGMukD8eY/s900-c-k-no-rj-c0xffffff/photo.jpg",
			  "width": "900",
			  "height": "900"
			 }
			],
			"person": [
			 {
			  "url": "http://www.youtube.com/user/Node"
			 },
			 {
			  "url": "https://plus.google.com/116748607363412317039"
			 },
			 {
			  "url": "http://www.youtube.com/user/Node"
			 },
			 {
			  "url": "https://plus.google.com/116748607363412317039"
			 }
			],
			"metatags": [
			 {
			  "theme-color": "#e62117",
			  "title": "Node",
			  "og:site_name": "YouTube",
			  "og:url": "https://www.youtube.com/user/Node",
			  "og:title": "Node",
			  "og:image": "https://yt3.ggpht.com/-217Xf1nmtEc/AAAAAAAAAAI/AAAAAAAAAAA/WCrGMukD8eY/s900-c-k-no-rj-c0xffffff/photo.jpg",
			  "og:description": "Sam, Niko, Brandon and D bringing you quality gaming videos as well as live action content every single week! For business inquires please contact us at, inf...",
			  "al:ios:app_store_id": "544007664",
			  "al:ios:app_name": "YouTube",
			  "al:ios:url": "vnd.youtube://user/UCI4Wh0EQPjGx2jJLjmTsFBQ",
			  "al:android:url": "https://www.youtube.com/user/Node?feature=applinks",
			  "al:android:app_name": "YouTube",
			  "al:android:package": "com.google.android.youtube",
			  "al:web:url": "https://www.youtube.com/user/Node?feature=applinks",
			  "og:type": "profile",
			  "og:video:tag": "FreddieW",
			  "fb:app_id": "87741124305",
			  "fb:profile_id": "node",
			  "twitter:card": "summary",
			  "twitter:site": "@youtube",
			  "twitter:url": "https://www.youtube.com/user/Node",
			  "twitter:title": "Node",
			  "twitter:description": "Sam, Niko, Brandon and D bringing you quality gaming videos as well as live action content every single week! For business inquires please contact us at, inf...",
			  "twitter:image": "https://yt3.ggpht.com/-217Xf1nmtEc/AAAAAAAAAAI/AAAAAAAAAAA/WCrGMukD8eY/s900-c-k-no-rj-c0xffffff/photo.jpg",
			  "twitter:app:name:iphone": "YouTube",
			  "twitter:app:id:iphone": "544007664",
			  "twitter:app:name:ipad": "YouTube",
			  "twitter:app:id:ipad": "544007664",
			  "twitter:app:url:iphone": "vnd.youtube://user/UCI4Wh0EQPjGx2jJLjmTsFBQ",
			  "twitter:app:url:ipad": "vnd.youtube://user/UCI4Wh0EQPjGx2jJLjmTsFBQ",
			  "twitter:app:name:googleplay": "YouTube",
			  "twitter:app:id:googleplay": "com.google.android.youtube",
			  "twitter:app:url:googleplay": "https://www.youtube.com/user/Node"
			 }
			],
			"cse_image": [
			 {
			  "src": "https://yt3.ggpht.com/-217Xf1nmtEc/AAAAAAAAAAI/AAAAAAAAAAA/WCrGMukD8eY/s900-c-k-no-rj-c0xffffff/photo.jpg"
			 }
			]
		   }
		  },
		  {
		   "kind": "customsearch#result",
		   "title": "Node (networking) - Wikipedia, the free encyclopedia",
		   "htmlTitle": "\u003cb\u003eNode\u003c/b\u003e (networking) - Wikipedia, the free encyclopedia",
		   "link": "https://en.wikipedia.org/wiki/Node_(networking)",
		   "displayLink": "en.wikipedia.org",
		   "snippet": "In communication networks, a node (Latin nodus, 'knot') is either a connection \npoint, a redistribution point, or a communication endpoint (e.g. data terminal ...",
		   "htmlSnippet": "In communication networks, a \u003cb\u003enode\u003c/b\u003e (Latin nodus, &#39;knot&#39;) is either a connection \u003cbr\u003e\npoint, a redistribution point, or a communication endpoint (e.g. data terminal&nbsp;...",
		   "cacheId": "Cw2OdIyUXAEJ",
		   "formattedUrl": "https://en.wikipedia.org/wiki/Node_(networking)",
		   "htmlFormattedUrl": "https://en.wikipedia.org/wiki/\u003cb\u003eNode\u003c/b\u003e_(networking)",
		   "pagemap": {
			"metatags": [
			 {
			  "referrer": "origin-when-cross-origin"
			 }
			]
		   }
		  },
		  {
		   "kind": "customsearch#result",
		   "title": "Download | Node.js",
		   "htmlTitle": "Download | \u003cb\u003eNode\u003c/b\u003e.js",
		   "link": "https://nodejs.org/en/download/",
		   "displayLink": "nodejs.org",
		   "snippet": "Downloads. Current version: v4.4.3. Download the Node.js source code or a pre-\nbuilt installer for your platform, and start developing today. LTS. Recommended ...",
		   "htmlSnippet": "Downloads. Current version: v4.4.3. Download the \u003cb\u003eNode\u003c/b\u003e.js source code or a pre-\u003cbr\u003e\nbuilt installer for your platform, and start developing today. LTS. Recommended&nbsp;...",
		   "cacheId": "oShuzBQnXcsJ",
		   "formattedUrl": "https://nodejs.org/en/download/",
		   "htmlFormattedUrl": "https://\u003cb\u003enode\u003c/b\u003ejs.org/en/download/",
		   "pagemap": {
			"metatags": [
			 {
			  "author": "Node.js Foundation",
			  "viewport": "width=device-width, initial-scale=1.0"
			 }
			]
		   }
		  },
		  {
		   "kind": "customsearch#result",
		   "title": "Node (computer science) - Wikipedia, the free encyclopedia",
		   "htmlTitle": "\u003cb\u003eNode\u003c/b\u003e (computer science) - Wikipedia, the free encyclopedia",
		   "link": "https://en.wikipedia.org/wiki/Node_(computer_science)",
		   "displayLink": "en.wikipedia.org",
		   "snippet": "A node is a basic unit used in computer science. Nodes are devices or data \npoints on a larger network. Devices such as a personal computer, cell phone, ...",
		   "htmlSnippet": "A \u003cb\u003enode\u003c/b\u003e is a basic unit used in computer science. \u003cb\u003eNodes\u003c/b\u003e are devices or data \u003cbr\u003e\npoints on a larger network. Devices such as a personal computer, cell phone,&nbsp;...",
		   "cacheId": "U8SoYJKw7usJ",
		   "formattedUrl": "https://en.wikipedia.org/wiki/Node_(computer_science)",
		   "htmlFormattedUrl": "https://en.wikipedia.org/wiki/\u003cb\u003eNode\u003c/b\u003e_(computer_science)",
		   "pagemap": {
			"cse_thumbnail": [
			 {
			  "width": "200",
			  "height": "132",
			  "src": "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQAMbc92dpQOLyOQqzDChn8Jf0oXGJxhjifwbyY8UG2RRORYSe7a91Nh9g"
			 }
			],
			"metatags": [
			 {
			  "referrer": "origin-when-cross-origin"
			 }
			],
			"cse_image": [
			 {
			  "src": "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5b/6n-graf.svg/250px-6n-graf.svg.png"
			 }
			]
		   }
		  },
		  {
		   "kind": "customsearch#result",
		   "title": "GitHub - nodejs/node: Node.js JavaScript runtime",
		   "htmlTitle": "GitHub - nodejs/\u003cb\u003enode\u003c/b\u003e: \u003cb\u003eNode\u003c/b\u003e.js JavaScript runtime",
		   "link": "https://github.com/nodejs/node",
		   "displayLink": "github.com",
		   "snippet": "node - Node.js JavaScript runtime :sparkles::turtle::rocket::sparkles:",
		   "htmlSnippet": "\u003cb\u003enode\u003c/b\u003e - \u003cb\u003eNode\u003c/b\u003e.js JavaScript runtime :sparkles::turtle::rocket::sparkles:",
		   "cacheId": "48Ua044jrKIJ",
		   "formattedUrl": "https://github.com/nodejs/node",
		   "htmlFormattedUrl": "https://github.com/\u003cb\u003enode\u003c/b\u003ejs/\u003cb\u003enode\u003c/b\u003e",
		   "pagemap": {
			"cse_thumbnail": [
			 {
			  "width": "225",
			  "height": "225",
			  "src": "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRT3rCwNQCmKFASdZIzXOBussVepo-5XMrAX_O4qOoPf3hm_nnHnU2me7ie"
			 }
			],
			"softwaresourcecode": [
			 {
			  "author": "nodejs",
			  "name": "node",
			  "about": "Node.js JavaScript runtime",
			  "url": "https://nodejs.org",
			  "keywords": "JavaScript",
			  "datemodified": "12 days ago",
			  "license": "LICENSE",
			  "text": "Node.js Node.js is a JavaScript runtime built on Chrome's V8 JavaScript engine. Node.js uses an event-driven, non-blocking I/O model that makes it lightweight and efficient. The Node.js package..."
			 }
			],
			"metatags": [
			 {
			  "viewport": "width=1020",
			  "fb:app_id": "1401488693436528",
			  "twitter:image:src": "https://avatars3.githubusercontent.com/u/9950313?v=3&s=400",
			  "twitter:site": "@github",
			  "twitter:card": "summary",
			  "twitter:title": "nodejs/node",
			  "twitter:description": "node - Node.js JavaScript runtime :sparkles::turtle::rocket::sparkles:",
			  "og:image": "https://avatars3.githubusercontent.com/u/9950313?v=3&s=400",
			  "og:site_name": "GitHub",
			  "og:type": "object",
			  "og:title": "nodejs/node",
			  "og:url": "https://github.com/nodejs/node",
			  "og:description": "node - Node.js JavaScript runtime :sparkles::turtle::rocket::sparkles:",
			  "browser-stats-url": "https://api.github.com/_private/browser/stats",
			  "browser-errors-url": "https://api.github.com/_private/browser/errors",
			  "pjax-timeout": "1000",
			  "msapplication-tileimage": "/windows-tile.png",
			  "msapplication-tilecolor": "#ffffff",
			  "google-analytics": "UA-3769691-2",
			  "octolytics-host": "collector.githubapp.com",
			  "octolytics-app-id": "github",
			  "octolytics-dimension-request_id": "42F94298:3F21:207C1:57292347",
			  "analytics-location": "/\u003cuser-name\u003e/\u003crepo-name\u003e",
			  "dimension1": "Logged Out",
			  "hostname": "github.com",
			  "expected-hostname": "github.com",
			  "js-proxy-site-detection-payload": "ODcyMjJjN2FjOTE1MTc5OTBlY2U3NjkwZDE4NTA1MWFiODk3MmY5ZGMzZGI4OWE2MjAyZDdjZTViYzNlM2ZiM3x7InJlbW90ZV9hZGRyZXNzIjoiNjYuMjQ5LjY2LjE1MiIsInJlcXVlc3RfaWQiOiI0MkY5NDI5ODozRjIxOjIwN0MxOjU3MjkyMzQ3IiwidGltZXN0YW1wIjoxNDYyMzEzNzk5fQ==",
			  "form-nonce": "4786915a0e60bc3175d7f1caaa43529927b568e3",
			  "go-import": "github.com/nodejs/node git https://github.com/nodejs/node.git",
			  "octolytics-dimension-user_id": "9950313",
			  "octolytics-dimension-user_login": "nodejs",
			  "octolytics-dimension-repository_id": "27193779",
			  "octolytics-dimension-repository_nwo": "nodejs/node",
			  "octolytics-dimension-repository_public": "true",
			  "octolytics-dimension-repository_is_fork": "false",
			  "octolytics-dimension-repository_network_root_id": "27193779",
			  "octolytics-dimension-repository_network_root_nwo": "nodejs/node"
			 }
			],
			"cse_image": [
			 {
			  "src": "https://avatars3.githubusercontent.com/u/9950313?v=3&s=400"
			 }
			],
			"listitem": [
			 {
			  "url": "Code",
			  "name": "Code",
			  "position": "1"
			 },
			 {
			  "url": "Issues 438",
			  "name": "Issues",
			  "position": "2"
			 },
			 {
			  "url": "Pull requests 204",
			  "name": "Pull requests",
			  "position": "3"
			 }
			]
		   }
		  },
		  {
		   "kind": "customsearch#result",
		   "title": "NODE Berlin Oslo — Graphic Design Studio",
		   "htmlTitle": "\u003cb\u003eNODE\u003c/b\u003e Berlin Oslo — Graphic Design Studio",
		   "link": "http://nodeberlin.com/",
		   "displayLink": "nodeberlin.com",
		   "snippet": "NODE has been contributing with exhibition design and we are currently working \non a richly documented, research-based catalogue written by the exhibition's ...",
		   "htmlSnippet": "\u003cb\u003eNODE\u003c/b\u003e has been contributing with exhibition design and we are currently working \u003cbr\u003e\non a richly documented, research-based catalogue written by the exhibition&#39;s&nbsp;...",
		   "cacheId": "7I2YERf65UoJ",
		   "formattedUrl": "nodeberlin.com/",
		   "htmlFormattedUrl": "\u003cb\u003enode\u003c/b\u003eberlin.com/",
		   "pagemap": {
			"cse_thumbnail": [
			 {
			  "width": "196",
			  "height": "257",
			  "src": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR_ZbN55c7mr_4BzDpUiwMpWG0PacGTdoetaKaCUqrBLTB0AREHy-K-RV6-"
			 }
			],
			"metatags": [
			 {
			  "author": "André Pahl, http://primeclub.org",
			  "og:type": "article",
			  "fb:admins": "541732749,1035952547,685776310,716663512",
			  "og:title": "NODE Berlin Oslo — Graphic Design Studio",
			  "og:site_name": "Node Berlin Oslo",
			  "viewport": "user-scalable=yes, minimum-scale=0.1, maximum-scale=2.0, width=1200",
			  "apple-mobile-web-app-capable": "yes",
			  "apple-mobile-web-app-status-bar-style": "black"
			 }
			],
			"cse_image": [
			 {
			  "src": "http://nodeberlin.com/_WORKBENCH/185/_792/CCA_AH_Cover_NODEBERLINOSLO_1_792.jpg"
			 }
			]
		   }
		  },
		  {
		   "kind": "customsearch#result",
		   "title": "Node (Java Platform SE 6)",
		   "htmlTitle": "\u003cb\u003eNode\u003c/b\u003e (Java Platform SE 6)",
		   "link": "http://docs.oracle.com/javase/6/docs/api/org/w3c/dom/Node.html",
		   "displayLink": "docs.oracle.com",
		   "snippet": "It represents a single node in the document tree. While all objects implementing \nthe Node interface expose methods for dealing with children, not all objects ...",
		   "htmlSnippet": "It represents a single \u003cb\u003enode\u003c/b\u003e in the document tree. While all objects implementing \u003cbr\u003e\nthe \u003cb\u003eNode\u003c/b\u003e interface expose methods for dealing with children, not all objects&nbsp;...",
		   "cacheId": "qRvJhMKBuTkJ",
		   "formattedUrl": "docs.oracle.com/javase/6/docs/api/org/w3c/dom/Node.html",
		   "htmlFormattedUrl": "docs.oracle.com/javase/6/docs/api/org/w3c/dom/\u003cb\u003eNode\u003c/b\u003e.html",
		   "pagemap": {
			"metatags": [
			 {
			  "date": "2015-11-19"
			 }
			]
		   }
		  },
		  {
		   "kind": "customsearch#result",
		   "title": "How To Node - NodeJS",
		   "htmlTitle": "How To \u003cb\u003eNode\u003c/b\u003e - NodeJS",
		   "link": "https://howtonode.org/",
		   "displayLink": "howtonode.org",
		   "snippet": "Community supported blog, teaches about the various tasks and fundamental \nconcepts to write effective code.",
		   "htmlSnippet": "Community supported blog, teaches about the various tasks and fundamental \u003cbr\u003e\nconcepts to write effective code.",
		   "cacheId": "SdOjhrztdHAJ",
		   "formattedUrl": "https://howtonode.org/",
		   "htmlFormattedUrl": "https://howto\u003cb\u003enode\u003c/b\u003e.org/",
		   "pagemap": {
			"metatags": [
			 {
			  "viewport": "width=device-width"
			 }
			]
		   }
		  },
		  {
		   "kind": "customsearch#result",
		   "title": "Node - Web APIs | MDN",
		   "htmlTitle": "\u003cb\u003eNode\u003c/b\u003e - Web APIs | MDN",
		   "link": "https://developer.mozilla.org/en-US/docs/Web/API/Node",
		   "displayLink": "developer.mozilla.org",
		   "snippet": "Mar 30, 2016 ... A Node is an interface from which a number of DOM types inherit, and allows \nthese various types to be treated (or tested) similarly.",
		   "htmlSnippet": "Mar 30, 2016 \u003cb\u003e...\u003c/b\u003e A \u003cb\u003eNode\u003c/b\u003e is an interface from which a number of DOM types inherit, and allows \u003cbr\u003e\nthese various types to be treated (or tested) similarly.",
		   "cacheId": "xuBjFFLO5c8J",
		   "formattedUrl": "https://developer.mozilla.org/en-US/docs/Web/API/Node",
		   "htmlFormattedUrl": "https://developer.mozilla.org/en-US/docs/Web/API/\u003cb\u003eNode\u003c/b\u003e",
		   "pagemap": {
			"metatags": [
			 {
			  "viewport": "width=device-width, initial-scale=1",
			  "og:type": "website",
			  "og:image": "https://developer.cdn.mozilla.net/static/img/opengraph-logo.dc4e08e2f6af.png",
			  "og:site_name": "Mozilla Developer Network",
			  "twitter:card": "summary",
			  "twitter:image": "https://developer.cdn.mozilla.net/static/img/opengraph-logo.dc4e08e2f6af.png",
			  "twitter:site": "@MozDevNet",
			  "twitter:creator": "@MozDevNet",
			  "og:title": "Node",
			  "og:url": "https://developer.mozilla.org/en-US/docs/Web/API/Node",
			  "twitter:url": "https://developer.mozilla.org/en-US/docs/Web/API/Node",
			  "twitter:title": "Node",
			  "og:description": "A Node is an interface from which a number of DOM types inherit, and allows these various types to be treated (or tested) similarly.",
			  "twitter:description": "A Node is an interface from which a number of DOM types inherit, and allows these various types to be treated (or tested) similarly."
			 }
			],
			"Breadcrumb": [
			 {
			  "title": "MDN",
			  "url": "MDN"
			 },
			 {
			  "title": "Web technology for developers",
			  "url": "Web technology for developers"
			 },
			 {
			  "title": "Web APIs",
			  "url": "Web APIs"
			 }
			],
			"cse_image": [
			 {
			  "src": "https://developer.cdn.mozilla.net/static/img/opengraph-logo.dc4e08e2f6af.png"
			 }
			]
		   }
		  }
		 ]
		}
	`)
	var result SearchResult
	err := json.Unmarshal(byt, &result)
	return err, result
}
