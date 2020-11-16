package client

import (
	"github.com/gin-gonic/gin"
	"go-tellme/internal/constants/model"
	"go-tellme/internal/utils"
	"go-tellme/platform/routers"
)

type Persistence interface {
	UniqueEmail(email string) error
	InsertAccount(user *model.User) error
	GetAccount(email string) (*model.User, error)
}

type Repository interface {
}

type service struct {
	userPersistence Persistence
	userRepository  Repository
	JWT             *utils.JWT
}

type Usecase interface {
	Register(user *model.User) (int, error)
	Login(email, password string) (int, error)

	GetToken(ID int) (string, error)
	VerifyToken(ctx *gin.Context) error
}

func InitializeDomain(persistence Persistence, repository Repository, jwt *utils.JWT) Usecase {
	return &service{persistence, repository, jwt}
}

type Handler interface {
	PageHome(ctx *gin.Context)
	PageLogin(ctx *gin.Context)
	PageLoginPerform(ctx *gin.Context)
	PageSignup(ctx *gin.Context)
	PageSignupPerform(ctx *gin.Context)
	PageAbout(ctx *gin.Context)
	PageLogout(ctx *gin.Context)

	EnsureIndex() gin.HandlerFunc
	EnsureLoggedIn() gin.HandlerFunc
	EnsureNotLoggedIn() gin.HandlerFunc
}

type Route interface {
	Routers() []*routers.Router
}
