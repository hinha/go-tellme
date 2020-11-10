package bot

import (
	"github.com/sirupsen/logrus"
	"go-tellme/internal/constants/model"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func (s *service) Register(user *tb.User) error {
	log := logrus.WithFields(logrus.Fields{
		"domain":  "telegram bot",
		"action":  "Authenticate",
		"usecase": "Register",
	})
	var firstTime bool
	err := s.botRepository.GetUserID(strconv.Itoa(user.ID))
	if err != nil {
		log.WithField("type", "GetUserID Repository").Errorln(err)
		firstTime = true
	}

	coll := &model.UserBot{
		UserID:       strconv.Itoa(user.ID),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		LanguageCode: user.LanguageCode,
		//CreatedAt: time.Unix(create, 0).String(),
		LogMessage: []model.LogMessage{{}},
		//Token: []model.Token{{}},
	}
	if firstTime {
		if err := s.botRepository.InsertUser(coll); err != nil {
			log.WithField("type", "InsertUser Repository").Errorln(err)
			return err
		}
		firstTime = false
	}

	return nil
}

func (s *service) GetUserByID(ID int) error {
	log := logrus.WithFields(logrus.Fields{
		"domain":  "telegram bot",
		"action":  "Authenticate",
		"usecase": "GetUserByID",
	})

	err := s.botRepository.GetUserID(strconv.Itoa(ID))
	if err != nil {
		log.WithField("type", "GetUserID Repository").Errorln(err)
		return err
	}

	return nil
}

func (s *service) GetInfoToken(ID int) (string, error) {
	log := logrus.WithFields(logrus.Fields{
		"domain":  "telegram bot",
		"action":  "Info Token",
		"usecase": "GetInfoToken",
	})

	token, err := s.botPersistence.InfoToken(strconv.Itoa(ID))
	if err != nil {
		log.WithField("type", "InfoToken persistence").Errorln(err)
		return "", err
	}

	return token, nil
}
