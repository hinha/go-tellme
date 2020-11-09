package bot

import (
	"errors"
	"go-tellme/internal/constants/model"
)

func (s *service) FindUserFirst(username string) (*model.UserBot, error) {
	user, err := s.botPersistence.FindUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) FindTokenFirst(username string) (string, error) {
	token, err := s.botPersistence.FindToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) GetKeyValidation(token string) error {
	if token != "abc" {
		return errors.New("test")
	}
	return nil
}

func (s *service) GetErrorToken(username string) error {
	err := s.botCaching.GetAction(username, "/token")
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetErrorStart(username string) error {
	err := s.botCaching.GetStartAction(username)
	if err != nil {
		return err
	}

	return nil
}
