package main

import (
	"github.com/coocood/freecache"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go-tellme/internal/glue"
	"go-tellme/internal/handler/rest"
	"go-tellme/internal/module/client"
	"go-tellme/internal/repository"
	"go-tellme/internal/storage/persistence"
	"go-tellme/internal/utils"
	"go-tellme/platform"
	"go-tellme/platform/routers"
	"log"
	"time"

	//"go-tellme/platform"
	"os"
)

const (
	host   = "localhost:5000"
	domain = "client"
	// 100 * 1024*1024 represents 100 Megabytes.
	cacheSize = 100 * 1024 * 1024
)

var (
	basePath string
	JwtKey   = os.Getenv("SECRET_KEY")
	ujwt     *utils.JWT
)

func init() {

	basePath, _ = os.Getwd()
	var filePath string

	filePath = basePath + "/cmd/user/.env"
	_ = godotenv.Load(filePath)

	if JwtKey == "" {
		JwtKey = "secret"
	}

	ujw, err := utils.New(&utils.JWT{
		Realm:         "test",
		Key:           []byte(JwtKey),
		Timeout:       time.Minute * 60,
		MaxRefresh:    time.Minute * 10,
		IdentityKey:   "id",
		TokenLookup:   "cookie: session",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		PayloadFunc: func(data interface{}) utils.MapClaims {
			claims := jwt.MapClaims{}
			params := data.(map[string]interface{})
			if len(params) > 0 {
				for k, v := range params {
					claims[k] = v
				}
			}
			return utils.MapClaims(claims)
		},
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	ujwt = ujw
}

func main() {
	// initialize connection
	dbConn := platform.ConnectionDatabase(os.Getenv("MONGO_URL"), os.Getenv("MONGO_DB"), domain)
	cache := freecache.NewCache(cacheSize)

	dbPersistence := persistence.ClientInit(dbConn)
	repo := repository.WebInit(dbPersistence)
	usecase := client.InitializeDomain(dbPersistence, repo, ujwt)
	handle := rest.NewHandlerWeb(usecase, cache)
	router := glue.WebInit(handle).Routers()

	serve := routers.Initialize(host, router, domain, "client", "secretCookie")
	serve.Serve()

}
