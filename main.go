package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var googleFavIconURL = "https://www.google.com/s2/favicons?domain="

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("Please provide site URL")
	}

	var siteURL string

	if !strings.HasPrefix(os.Args[1], "http") {
		siteURL = "http://" + os.Args[1]
	} else {
		siteURL = os.Args[1]
	}

	checkedURL, err := url.ParseRequestURI(siteURL)
	if err != nil {
		log.Fatal(err)
	}

	faviconFile, err := os.OpenFile(checkedURL.Host+".png", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer faviconFile.Close()

	response, err := http.Get(googleFavIconURL + checkedURL.Host)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		_, err := io.Copy(faviconFile, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
