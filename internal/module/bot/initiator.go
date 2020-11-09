package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go-tellme/internal/constants/model"
	"go-tellme/platform/grpc/gen"
)

type Persistence interface {
	CreateUser(user *model.UserBot) error
	FindUsername(username string) (*model.UserBot, error)
	FindToken(username string) (string, error)
}

type Caching interface {
	SaveToken(username, token string) error
	SaveAction(username, action string) error
	GetToken(username string) (string, error)
	GetStartAction(username string) error
	GetAction(username, action string) error
}

type Repository interface {
	IsUser(token string) error
	ChatBot(payload *gen.ChatPayload) (*gen.ChatResponse, error)
}

type service struct {
	botPersistence Persistence
	botCaching     Caching
	botRepository  Repository
}

type Usecase interface {
	CreateUserFirst(c *tgbotapi.Message) error
	FindUserFirst(username string) (*model.UserBot, error)
	FindTokenFirst(username string) (string, error)

	Action(username, action string) error
	GetInputToken(username string) (string, error)
	GetErrorStart(username string) error // validasi untuk melanjutkan chatbot
	GetErrorToken(username string) error
	GetKeyValidation(token string) error
	CommandsBot(payload *gen.ChatPayload) (string, error)
	InsertToken(username, token string) error
}

func InitializeDomain(persistence Persistence, caching Caching, repository Repository) Usecase {
	return &service{
		botPersistence: persistence,
		botCaching:     caching,
		botRepository:  repository,
	}
}

type Handler interface {
	Bot(c *tgbotapi.Message)
	FirstMessage(c *tgbotapi.Message)
	Help(c *tgbotapi.Message)
	Greeting(c *tgbotapi.Message)
	InfoToken(c *tgbotapi.Message)
	ServeBot()
	Commands(c tgbotapi.Update) error
}
