package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"go-tellme/internal/module/bot"
	"go-tellme/platform/grpc/gen"
	"log"
	"net/url"
	"strconv"
	"time"
)

const (
	msgError     = "Tidak dapat melanjutkan silahkan ketik /start"
	msgToken     = "Tidak dapat melanjutkan silahkan ketik /token"
	errorService = "Bot sedang offline"
)

type botService struct {
	usecase      bot.Usecase
	domain       string
	secretKey    string
	ReplyMessage string
}

func (b *botService) FirstMessage(c *tgbotapi.Message) {
	user, err := b.usecase.FindUserFirst(c.From.UserName)
	if err != nil {
		b.ReplyMessage = msgError
		return
	}
	b.Greeting(c)

	if user.Token == "" {

	}
	//b.InfoToken(c)
	//b.Bot(c)
}

func (b *botService) InfoToken(c *tgbotapi.Message) {
	token, err := b.usecase.FindTokenFirst(c.From.UserName)
	if err != nil {
		b.ReplyMessage = errorService
		return
	}
	b.ReplyMessage = "Your Token: " + token
	return
}

func (b *botService) Bot(c *tgbotapi.Message) {
	token, err := b.usecase.FindTokenFirst(c.From.UserName)
	if err != nil {
		b.ReplyMessage = errorService
		return
	}
	if token == "" {
		b.ReplyMessage = "Please input your token"
		return
	}

	offMsg := "Maaf noel sedang offline coba lagi nanti ya."

	msg, err := b.usecase.CommandsBot(&gen.ChatPayload{
		Message:  c.Text,
		UserId:   strconv.Itoa(c.MessageID),
		CreateAt: time.Now().String(),
	})

	if err != nil {
		b.ReplyMessage = offMsg
		return
	}

	decodeMessage, err := url.QueryUnescape(msg)
	if err != nil {
		b.ReplyMessage = offMsg
		return
	}
	b.ReplyMessage = decodeMessage
	return
}

func (b *botService) Greeting(c *tgbotapi.Message) {
	_ = b.usecase.Action(c.From.UserName, c.Text)
	var text string
	tm := time.Now()

	switch {
	case tm.Hour() < 10:
		text = "Pagi"
	case tm.Hour() < 15:
		text = "Siang"
	case tm.Hour() < 18:
		text = "Sore"
	default:
		text = "Malam"
	}
	_, err := b.usecase.FindUserFirst(c.From.UserName)
	if err != nil {
		err = b.usecase.CreateUserFirst(c)
		if err != nil {
			b.ReplyMessage = errorService
			return
		}
	}
	b.ReplyMessage = "Halo Selamat " + text
}

func (b *botService) Help(c *tgbotapi.Message) {
	b.ReplyMessage = "I can help you." + "You can control me by sending these commands: \n\n /token - set <BOT_TOKEN>"
}

func (b *botService) ServeBot() {
	conn, err := tgbotapi.NewBotAPI(b.secretKey)
	if err != nil {
		log.Fatal(err)
	}

	logrus.WithFields(logrus.Fields{
		"domain":  b.domain,
		"account": conn.Self.UserName,
	}).Info("Starts Serving Telegram")

	uu := tgbotapi.NewUpdate(0)
	uu.Timeout = 60

	updates, err := conn.GetUpdatesChan(uu)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		err := b.Commands(update)
		if err != nil {
			break
		}
		if update.Message == nil {
			continue
		}

		logrus.Info(fmt.Printf("[%s] %s \n", update.Message.From.UserName, update.Message.Text))

		_, err = conn.Send(tgbotapi.NewMessage(update.Message.Chat.ID, b.ReplyMessage))
		if err != nil {
			fmt.Println("error: ", err)
		}
	}
}

// Basic Logic
// /start -> Greeting
//	[1] /token
// /help -> Help
func (b *botService) Commands(c tgbotapi.Update) error {

	msg := c.Message.Text
	if msg == "/start" {
		b.Greeting(c.Message)
		return nil
	} else if msg == "/help" {
		b.Help(c.Message)
		return nil
	} else if msg == "/token" {
		b.InfoToken(c.Message)
		return nil
	} else {
		b.FirstMessage(c.Message)
	}

	return nil
}

func NewHandlerBot(usecase bot.Usecase, domain, key string) bot.Handler {
	return &botService{usecase: usecase, domain: domain, secretKey: key}
}
