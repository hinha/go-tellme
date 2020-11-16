package client

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-tellme/internal/utils"
)

func (s *service) Login(email, password string) (int, error) {
	log := logrus.WithFields(logrus.Fields{
		"domain":  "user",
		"action":  "User Login",
		"usecase": "Login",
	})

	user, err := s.userPersistence.GetAccount(email)
	if err != nil {
		log.WithField("type", "GetAccount persistence").Errorln(err)
		return 0, err
	}

	compare := utils.CheckPasswordHash(password, user.Password)
	if compare {
		return user.ID, nil
	} else {
		log.WithField("type", "CheckPasswordHash utils").Errorln("password")
		return 0, errors.New("password not match")
	}
}

func (s *service) VerifyToken(ctx *gin.Context) error {
	log := logrus.WithFields(logrus.Fields{
		"domain":  "user",
		"action":  "User Login",
		"usecase": "Verify Token",
	})

	_, err := s.JWT.CheckTokenExpire(ctx)
	if err != nil {
		log.WithField("type", "CheckTokenExpire JWT").Errorln(err)
		return err
	}

	return nil
}
