package platform

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type ConnectionsTelegram interface {
	OpenTG() (*tgbotapi.BotAPI, tgbotapi.UpdateConfig)
}

type connectionStringTelegram struct {
	client string
	domain string
}

func InitializeTelegramBot(client, domain string) ConnectionsTelegram {
	return &connectionStringTelegram{client: client, domain: domain}
}

func (c connectionStringTelegram) OpenTG() (*tgbotapi.BotAPI, tgbotapi.UpdateConfig) {
	bot, err := tgbotapi.NewBotAPI(c.client)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	uu := tgbotapi.NewUpdate(0)
	uu.Timeout = 60

	return bot, uu
}
