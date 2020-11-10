package tele

import (
	"fmt"
	"go-tellme/internal/constants/button"
	"go-tellme/internal/module/bot"
	tb "gopkg.in/tucnak/telebot.v2"
	"regexp"
	"strings"
)

const (
	msgError      = "Tidak dapat melanjutkan silahkan ketik /start"
	msgToken      = "Tidak dapat melanjutkan silahkan ketik /token"
	errorService  = "Bot sedang offline"
	cmdTokenEmpty = "Ketik perintah berikut: /token <TOKEN>"
)

var messageCallback string

type botService struct {
	usecase bot.Usecase
	client  *tb.Bot
}

func (b *botService) Info(m *tb.Message) {
	messageCallback = "Info"
	if err := b.usecase.GetUserByID(m.Sender.ID); err != nil {
		messageCallback = "Ops something went wrong or bot error"
	}

	b.client.Send(m.Sender, messageCallback)
}

func (b *botService) Start(m *tb.Message) {
	button.Menu.Reply(
		button.Menu.Row(button.BtnInfo),
		button.Menu.Row(button.BtnSettings),
		button.Menu.Row(button.BtnToken),
		button.Menu.Row(button.BtnHelp),
	)
	button.Selector.Inline(
		button.Selector.Row(button.BtnPrev, button.BtnNext),
	)

	messageCallback = "Welcome"
	if err := b.usecase.Register(m.Sender); err != nil {
		messageCallback = "Ops something went wrong or bot error"
	}

	b.client.Send(m.Sender, messageCallback, button.Menu)
}

func (b *botService) Token(m *tb.Message) {

	spCommand := strings.Split(m.Text, " ")

	if len(spCommand) == 1 {
		b.client.Send(m.Sender, "insert token with argument: /token <YOUR_TOKEN>")
		return
	} else if len(spCommand) == 2 {
		if spCommand[0] == "/token" {
			b.client.Send(m.Sender, "Successfully insert token")
			return
		}
	}

}

func (b *botService) CommandAI(m *tb.Message) {
	prefix := "^/([a-zA-Z0-9_-]{1,64})"
	r, _ := regexp.Compile(prefix)
	got := r.MatchString(m.Text)

	if !got {
		fmt.Println(m.Text)
		b.client.Send(m.Sender, "Command AI")
	}

}

func (b *botService) InfoButton(m *tb.Message) {
	b.client.Send(m.Sender, "Info button")
}

func (b *botService) TokenButton(m *tb.Message) {

	token, err := b.usecase.GetInfoToken(m.Sender.ID)
	if err != nil {
		b.client.Send(m.Sender, errorService)
		return
	}

	if token == "" {
		b.client.Send(m.Sender, "Token: empty")
		b.client.Send(m.Sender, cmdTokenEmpty)
		return
	}

	b.client.Send(m.Sender, fmt.Sprintf("Token: %s", token))
}

func (b *botService) HelpButton(m *tb.Message) {
	b.client.Send(m.Sender, "Help button")
}

func NewHandlerBot(usecase bot.Usecase, client *tb.Bot) bot.Handler {
	return &botService{usecase: usecase, client: client}
}
