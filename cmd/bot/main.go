package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const token = "463334009:AAEHiO4nel30tyNEKgNMhZtqzn5HnoZHHHU"

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//	u := tgbotapi.NewUpdate(0)
	//	u.Timeout = 60
	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Command() == "help" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "HELP command manual")
				bot.Send(msg)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You sent: "+update.Message.Text)

			bot.Send(msg)
		}
	}
}
