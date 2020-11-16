package client

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-tellme/internal/constants/model"
	"go-tellme/internal/utils"
	"math/rand"
	"strconv"
	"time"
)

func (s *service) Register(user *model.User) (int, error) {
	var isUser bool
	log := logrus.WithFields(logrus.Fields{
		"domain":  "user",
		"action":  "User Register",
		"usecase": "Register",
	})

	if err := s.userPersistence.UniqueEmail(user.Email); err != nil {
		isUser = true
	}

	years, months, days := time.Now().Date()
	if isUser {

		hasher, err := utils.HashPassword(user.Password)
		if err != nil {
			log.WithField("type", "HashPassword utils").Errorln(err)
			return 0, err
		}

		id, _ := strconv.Atoi(fmt.Sprintf("%d%d%d%d", int(months), years, randomInt(100001, 9999999), days))
		user.ID = id
		user.AccessKey = []model.Keys{}
		user.Password = hasher

		err = s.userPersistence.InsertAccount(user)
		if err != nil {
			log.WithField("type", "InsertAccount persistence").Errorln(err)
			return 0, err
		}

		return user.ID, nil
	} else {
		log.WithField("type", "InsertAccount persistence").Errorln("must unique email")
		return 0, errors.New("must unique email")
	}
}

func (s *service) GetToken(ID int) (string, error) {
	log := logrus.WithFields(logrus.Fields{
		"domain":  "user",
		"action":  "Token",
		"usecase": "GetToken",
	})

	data := map[string]interface{}{
		"user_id": ID,
	}

	token, _, err := s.JWT.TokenGenerator(data)
	if err != nil {
		log.WithField("type", "TokenGenerator utils").Errorln(err)
		return "", err
	}

	return token, nil
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
