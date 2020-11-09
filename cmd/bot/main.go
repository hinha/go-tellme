package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	gcache "github.com/patrickmn/go-cache"
	"go-tellme/internal/handler"
	"go-tellme/internal/module/bot"
	"go-tellme/internal/repository"
	"go-tellme/internal/storage/cache"
	"go-tellme/internal/storage/persistence"
	"go-tellme/platform"
	"go-tellme/platform/grpc"
	"log"
	"os"
	"time"
)

const (
	domain     = "telegram"
	ApiKeyTele = "1312609963:AAGiUpARgdhlMonxUjLGo-7piuUTl40wuVc"
)

var basePath string

func init() {
	basePath, _ = os.Getwd()
	var filePath string

	if len(os.Args) == 0 {
		panic("environment variable required")
	}

	switch os.Args[1] {
	case "development":
		filePath = basePath + "/cmd/bot/.env.development"
	default:
		filePath = basePath + "/cmd/bot/.env"
	}
	if err := godotenv.Load(filePath); err != nil {
		log.Fatal("Error loading .env file")
	}

	_ = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("1"),
			tgbotapi.NewKeyboardButton("2"),
			tgbotapi.NewKeyboardButton("3"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("4"),
			tgbotapi.NewKeyboardButton("5"),
			tgbotapi.NewKeyboardButton("6"),
		),
	)
}

func main() {
	// initialize connection
	dbConn := platform.ConnectionDatabase(os.Getenv("MONGO_URL"), os.Getenv("MONGO_DB"))

	// Configure amqp connection
	amqpString := fmt.Sprintf("amqp://%s:%s@%s/%s",
		os.Getenv("AMQP_USER"),
		os.Getenv("AMQP_PASS"),
		os.Getenv("AMQP_HOST"),
		os.Getenv("AMQP_V"))
	amqpClient := platform.InitializeAmqp(amqpString, domain)
	amqpConn := amqpClient.Open()

	// Configure grpc connection
	rpcClient := grpc.InitializeGrpc(os.Getenv("RPC_HOST"), os.Getenv("RPC_PORT"), domain)
	rpcConn := rpcClient.Open()
	defer rpcConn.Close()

	gCache := cache.InitCache(gcache.New(10*time.Minute, 24*time.Hour))

	dbPersistence := persistence.BotInit(dbConn)
	repo := repository.BotInit(dbPersistence, amqpConn, rpcConn)
	usecase := bot.InitializeDomain(dbPersistence, gCache, repo)
	handle := handler.NewHandlerBot(usecase, domain, ApiKeyTele)
	handle.ServeBot()
}
