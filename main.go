package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)


func main() {
	token := os.Getenv("TG_WEXORTBOT_TOKEN")

	if token == "" {
		log.Fatal("no token found")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("create botAPI failed", err)
		return
	}

	log.Printf("botAPI init: %s", bot.Self.UserName)

	config := tgbotapi.UpdateConfig{}
	updates, err := bot.GetUpdatesChan(config)

	for update := range updates {
		if strings.Contains(update.Message.Text, "学德语") {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Fuck off")
			_, err = bot.Send(msg)
			if err != nil {
				log.Println("Unable to send message:", err)
				return
		}

		}
	}
}
