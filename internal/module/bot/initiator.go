package bot

import (
	"go-tellme/internal/constants/model"
	"go-tellme/platform/grpc/gen"
	"go-tellme/platform/routers"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Persistence interface {
	GetUserID(ID string) (*model.UserBot, error)
	InsertUser(user *model.UserBot) (*model.UserBot, error)
	InfoToken(ID string) (string, error)

	//CreateUser(client *model.UserBot) error
	//FindUsername(username string) (*model.UserBot, error)
	//FindToken(username string) (string, error)
}

type Caching interface {
	SaveID(ID string, user *model.UserBot) error
}

type Repository interface {
	GetUserID(ID string) error
	InsertUser(user *model.UserBot) error

	IsUser(token string) error
	ChatBot(payload *gen.ChatPayload) (*gen.ChatResponse, error)
}

type service struct {
	botPersistence Persistence
	botCaching     Caching
	botRepository  Repository
}

type Usecase interface {
	Register(user *tb.User) error
	GetUserByID(ID int) error
	GetInfoToken(ID int) (string, error)

	//CreateUserFirst(c *tgbotapi.Message) error
	//FindUserFirst(username string) (*model.UserBot, error)
	//FindTokenFirst(username string) (string, error)
	//
	//Action(username, action string) error
	//GetInputToken(username string) (string, error)
	//GetErrorStart(username string) error // validasi untuk melanjutkan chatbot
	//GetErrorToken(username string) error
	//GetKeyValidation(token string) error
	//CommandsBot(payload *gen.ChatPayload) (string, error)
	//InsertToken(username, token string) error
}

func InitializeDomain(persistence Persistence, caching Caching, repository Repository) Usecase {
	return &service{
		botPersistence: persistence,
		botCaching:     caching,
		botRepository:  repository,
	}
}

type Handler interface {
	Start(m *tb.Message)
	Info(m *tb.Message)
	Token(m *tb.Message)
	CommandAI(m *tb.Message)

	// Button
	InfoButton(m *tb.Message)
	TokenButton(m *tb.Message)
	HelpButton(m *tb.Message)

	//Bot(c *tgbotapi.Message)
	//FirstMessage(c *tgbotapi.Message)
	//Help(c *tgbotapi.Message)
	//Greeting(c *tgbotapi.Message)
	//InfoToken(c *tgbotapi.Message)
	//ServeBot()
	//Commands(c tgbotapi.Update) error
}

type Route interface {
	Routers() []*routers.RouterBot
}
