package persistence

import (
	"github.com/Kamva/mgm/v3"
	"github.com/bradfitz/slice"
	"go-tellme/internal/constants/model"
	"go-tellme/internal/module/bot"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type botPersistence struct {
	db *mongo.Database
}

func (p *botPersistence) GetUserID(ID string) (*model.UserBot, error) {
	user := new(model.UserBot)
	coll := mgm.Coll(user)
	if err := coll.First(bson.M{"user_id": ID}, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (p *botPersistence) InsertUser(user *model.UserBot) (*model.UserBot, error) {
	coll := mgm.Coll(user)
	if err := coll.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (p *botPersistence) InfoToken(ID string) (string, error) {
	var token string
	user := new(model.UserBot)
	coll := mgm.Coll(user)
	if err := coll.First(bson.M{"user_id": ID}, user); err != nil {
		return "", err
	}

	if len(user.Token) == 0 {
		token = ""
	} else {
		slice.Sort(user.Token[:], func(i, j int) bool {
			a, _ := strconv.Atoi(user.Token[i].CreatedAt)
			b, _ := strconv.Atoi(user.Token[j].CreatedAt)
			return a > b
		})

		token = user.Token[0].Key
	}

	return token, nil
}

//func (p *botPersistence) CreateUser(user *model.UserBot) error {
//	coll := mgm.Coll(user)
//	if err := coll.Create(user); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (p *botPersistence) FindToken(username string) (string, error) {
//	user := new(model.UserBot)
//	coll := mgm.Coll(user)
//	if err := coll.First(bson.M{"username": username}, user); err != nil {
//		return "", err
//	}
//
//	return user.Token, nil
//}
//
//func (p *botPersistence) FindUsername(username string) (*model.UserBot, error) {
//	user := new(model.UserBot)
//	coll := mgm.Coll(user)
//	if err := coll.First(bson.M{"username": username}, user); err != nil {
//		return nil, err
//	}
//
//	return user, nil
//}

func BotInit(db *mongo.Database) bot.Persistence {
	return &botPersistence{db: db}
}
