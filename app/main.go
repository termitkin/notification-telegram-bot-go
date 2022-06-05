package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetReq(url string) {
	res, err := http.Get(url)

	if err == nil {
		defer res.Body.Close()
	} else {
		fmt.Println("Get request not sent")
		fmt.Printf("URL: %s", url)
		fmt.Println(err)
	}
}

func BuildQuery(message string) string {
	TelegramBotChatId := os.Getenv("TELEGRAM_BOT_CHAT_ID")

	if len(TelegramBotChatId) == 0 {
		log.Fatal("Telegram bot chat id not set")
	}

	return fmt.Sprintf("chat_id=%s&text=%s", url.QueryEscape(TelegramBotChatId), url.QueryEscape(message))
}

func BuildTelegramAPIUrl(query string) string {
	TelegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	if len(TelegramBotToken) == 0 {
		log.Fatal("Telegram bot token not set")
	}

	return fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?%s", TelegramBotToken, query)
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
	}

	text := string(body)

	if len(text) > 0 {
		query := BuildQuery(text)
		url := BuildTelegramAPIUrl(query)

		GetReq(url)

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
