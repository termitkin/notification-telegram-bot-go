package url

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func GetUrlQuery(message string) string {
	TelegramBotChatId := os.Getenv("TELEGRAM_BOT_CHAT_ID")

	if len(TelegramBotChatId) == 0 {
		log.Fatal("Telegram bot chat id not set")
	}

	query := url.Values{}
	query.Add("chat_id", TelegramBotChatId)
	query.Add("text", message)

	return query.Encode()
}

func GetUrl(query string) string {
	TelegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	if len(TelegramBotToken) == 0 {
		log.Fatal("Telegram bot token not set")
	}

	return fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?%s", TelegramBotToken, query)
}
