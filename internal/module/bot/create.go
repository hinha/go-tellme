package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go-tellme/internal/constants/model"
	"strconv"
	"time"
)

func (s *service) CreateUserFirst(c *tgbotapi.Message) error {
	//create := time.Now().Unix()
	i, err := strconv.ParseInt(strconv.Itoa(c.Date), 10, 64)
	if err != nil { // default
		i = 1603299612
	}
	tm := time.Unix(i, 0)

	err = s.botPersistence.CreateUser(&model.UserBot{
		UserID:       strconv.Itoa(c.From.ID),
		FirstName:    c.From.FirstName,
		LastName:     c.From.LastName,
		Username:     c.From.UserName,
		LanguageCode: c.From.LanguageCode,
		//CreatedAt: time.Unix(create, 0).String(),
		LogMessage: []model.LogMessage{{
			MessageID: c.MessageID,
			Text:      c.Text,
			CreatedAt: tm.String(),
		}},
		Token: "",
	})

	if err != nil {
		return err
	}
	return nil
}
