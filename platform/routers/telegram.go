package routers

import (
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

type RouterBot struct {
	URL        string
	URLBtn     tb.Btn
	Handler    func(m *tb.Message)
	HandlerBtn func(m *tb.Message)
}

type RoutingBot struct {
	domain  string
	service *tb.Bot
}

type Bot interface {
	Serve()
}

var handlersBot []*RouterBot

func InitializeBot(routers []*RouterBot, domain string, service *tb.Bot) Bot {
	handlersBot = routers
	return &RoutingBot{domain: domain, service: service}
}

func (u *RoutingBot) Serve() {

	for _, router := range handlersBot {
		if router.Handler != nil {
			u.service.Handle(router.URL, router.Handler)
		}
		if router.HandlerBtn != nil {
			u.service.Handle(&router.URLBtn, router.HandlerBtn)
		}
	}

	logrus.WithFields(logrus.Fields{
		"domain": u.domain,
	}).Info("Starts Serving Bot")

	u.service.Start()
}
