package glue

import (
	"go-tellme/internal/constants/button"
	"go-tellme/internal/module/bot"
	"go-tellme/platform/routers"
	tb "gopkg.in/tucnak/telebot.v2"
)

type telegramHandler struct {
	handler bot.Handler
}

func (t *telegramHandler) Routers() []*routers.RouterBot {
	return []*routers.RouterBot{
		{
			URL:     "/start",
			Handler: t.handler.Start,
		},
		{
			URL:     "/info",
			Handler: t.handler.Info,
		},
		{
			URL:     "/token",
			Handler: t.handler.Token,
		},
		{
			URL:     tb.OnText,
			Handler: t.handler.CommandAI,
		},
		{
			URLBtn:     button.BtnInfo,
			HandlerBtn: t.handler.InfoButton,
		},
		{
			URLBtn:     button.BtnToken,
			HandlerBtn: t.handler.TokenButton,
		},
		{
			URLBtn:     button.BtnHelp,
			HandlerBtn: t.handler.HelpButton,
		},
	}
}

func InitializeTelegram(handle bot.Handler) bot.Route {
	return &telegramHandler{handler: handle}
}
