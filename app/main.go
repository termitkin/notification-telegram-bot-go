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

func getUrlQuery(message string) string {
	TelegramBotChatId := os.Getenv("TELEGRAM_BOT_CHAT_ID")

	if len(TelegramBotChatId) == 0 {
		log.Fatal("Telegram bot chat id not set")
	}

	query := url.Values{}
	query.Add("chat_id", TelegramBotChatId)
	query.Add("text", message)

	return query.Encode()
}

func getUrl(query string) string {
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

		return
	}

	text := string(body)

	if len(text) > 0 {
		query := getUrlQuery(text)
		url := getUrl(query)

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
