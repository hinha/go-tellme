package bot

func (s *service) InsertToken(username, token string) error {
	_ = s.botCaching.SaveAction(username, token)
	return nil
}

func (s *service) GetInputToken(username string) (string, error) {
	token, err := s.botCaching.GetToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}
