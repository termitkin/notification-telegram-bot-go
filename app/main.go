package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/notification-telegram-bot-go/app/url"
)

func sendMessage(url string) {
	res, err := http.Get(url)

	if err == nil {
		defer res.Body.Close()
	} else {
		fmt.Printf("Get request not sent\n")
		fmt.Printf("URL: %s\n", url)
		fmt.Println(err)
	}
}

func requestHandler(res http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "text/plain") {
		fmt.Print(contentType)

		return
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println("Body not read")

		return
	}

	text := string(body)

	if len(text) > 0 {
		query := url.GetUrlQuery(text)
		url := url.GetUrl(query)

		sendMessage(url)

		_, err := res.Write([]byte("ok"))

		if err != nil {
			fmt.Println("Response not sent")
		}
	} else {
		fmt.Println("Body is empty")
	}
}

func main() {
	http.HandleFunc("/", requestHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
