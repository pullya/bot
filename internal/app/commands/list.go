package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pullya/bot/internal/service/product"
)

func (c *Commander) List(inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMsgText := "Here all the products: \n\n"
	products := productService.List()

	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	c.bot.Send(msg)
}
