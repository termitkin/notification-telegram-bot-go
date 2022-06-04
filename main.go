package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetReq(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Get request not sent")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()

		if err != nil {
			fmt.Println("Body not closed")
		}
	}(res.Body)
}

func BuildQuery(message string) string {
	TelegramBotChatId := os.Getenv("TELEGRAM_BOT_CHAT_ID")

	if len(TelegramBotChatId) == 0 {
		log.Fatal("Telegram bot chat id not set")
	}

	return fmt.Sprintf("chat_id=%s&text=%s", TelegramBotChatId, message)
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

	if contentType != "text/plain" {
		fmt.Println(contentType)

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
