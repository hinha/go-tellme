package model

import "github.com/Kamva/mgm/v3"

type UserTest struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"fist_name" bson:"fist_name"`
	Pages            int    `json:"pages" bson:"pages"`
}

type User struct {
	mgm.DefaultModel `bson:",inline"`
	ID               int    `json:"id"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	AccessKey        []Keys `json:"access_key"`
}

type Keys struct {
	Name  string  `json:"name"`
	Token []Token `json:"token"`
}
