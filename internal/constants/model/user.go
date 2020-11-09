package model

import "github.com/Kamva/mgm/v3"

type UserTest struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"fist_name" bson:"fist_name"`
	Pages            int    `json:"pages" bson:"pages"`
}
