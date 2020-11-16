package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/joho/godotenv"
	_ "github.com/patrickmn/go-cache"
	"go-tellme/internal/glue"
	"go-tellme/internal/handler/tele"
	"go-tellme/internal/module/bot"
	"go-tellme/internal/repository"
	"go-tellme/internal/storage/cache"
	"go-tellme/internal/storage/persistence"
	"go-tellme/platform"
	"go-tellme/platform/routers"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

const (
	domain = "telegram"
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

}

func main() {
	// initialize connection
	dbConn := platform.ConnectionDatabase(os.Getenv("MONGO_URL"), os.Getenv("MONGO_DB"), domain)

	// Configure amqp connection
	amqpString := fmt.Sprintf("amqp://%s:%s@%s/%s",
		os.Getenv("AMQP_USER"),
		os.Getenv("AMQP_PASS"),
		os.Getenv("AMQP_HOST"),
		os.Getenv("AMQP_V"))
	amqpClient := platform.InitializeAmqp(amqpString, domain)
	amqpConn := amqpClient.Open()

	// Configure grpc connection
	//rpcClient := grpc.InitializeGrpc(os.Getenv("RPC_HOST"), os.Getenv("RPC_PORT"), domain)
	//rpcConn := rpcClient.Open()
	//defer rpcConn.Close()

	// Configure mem-cache
	mc := memcache.New("0.0.0.0:11211")
	gCache := cache.InitCache(mc)

	dbPersistence := persistence.BotInit(dbConn)
	repo := repository.BotInit(dbPersistence, gCache, amqpConn, nil)
	usecase := bot.InitializeDomain(dbPersistence, gCache, repo)

	settings := tb.Settings{
		Token:  os.Getenv("TELEGRAM_KEY"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	newBot, err := tb.NewBot(settings)
	if err != nil {
		panic(err)
	}

	handle := tele.NewHandlerBot(usecase, newBot)
	router := glue.InitializeTelegram(handle).Routers()

	serve := routers.InitializeBot(router, domain, newBot)
	serve.Serve()
}
