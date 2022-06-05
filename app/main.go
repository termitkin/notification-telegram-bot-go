package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/termitkin/notification-telegram-bot-go/app/message"
	"github.com/termitkin/notification-telegram-bot-go/app/url"
)

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

		message.SendMessage(url)

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
