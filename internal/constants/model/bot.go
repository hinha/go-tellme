package model

import "github.com/Kamva/mgm/v3"

type UserBot struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           string       `json:"user_id" bson:"user_id"`
	FirstName        string       `json:"first_name" bson:"first_name"`
	LastName         string       `json:"last_name" bson:"last_name"`
	Username         string       `json:"username" bson:"username"`
	LanguageCode     string       `json:"language_code" bson:"language_code"`
	LogMessage       []LogMessage `json:"log_message" bson:"log_message"`
	Token            string       `json:"token" bson:"token"`
}

type LogMessage struct {
	MessageID int    `json:"message_id" bson:"message_id"`
	Text      string `json:"text" bson:"text"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}
