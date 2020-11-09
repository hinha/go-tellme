package bot

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/sirupsen/logrus"
	"go-tellme/platform/grpc/gen"
)

func (s *service) Action(username, action string) error {
	err := validation.Validate(action,
		validation.Required,
		validation.Length(1, 100),
		is.Alphanumeric,
	)
	if err != nil {
		_ = s.botCaching.SaveAction(username, action)
	}
	return nil
}

func (s *service) CommandsBot(payload *gen.ChatPayload) (string, error) {
	log := logrus.WithFields(logrus.Fields{
		"domain":  "telegram",
		"usecase": "CommandsBot",
	})

	err := s.botRepository.IsUser(payload.Message)
	if err != nil {
		log.WithField("type", "IsUser").Errorln(err)
	}

	ss, err := s.botRepository.ChatBot(payload)
	if err != nil {
		log.WithField("type", "ChatBot").Errorln(err)
		return "", err
	}

	if ss.Label == "twitter-trend" {
		ss.Message = "Noel tidak tahu keyword seperti apa. buat keyword dulu ya"
	}

	return ss.Message, nil
}
