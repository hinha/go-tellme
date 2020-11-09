package persistence

import (
	"github.com/Kamva/mgm/v3"
	"go-tellme/internal/constants/model"
	"go-tellme/internal/module/bot"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type botPersistence struct {
	db *mongo.Database
}

func (p *botPersistence) CreateUser(user *model.UserBot) error {
	coll := mgm.Coll(user)
	if err := coll.Create(user); err != nil {
		return err
	}

	return nil
}

func (p *botPersistence) FindToken(username string) (string, error) {
	user := new(model.UserBot)
	coll := mgm.Coll(user)
	if err := coll.First(bson.M{"username": username}, user); err != nil {
		return "", err
	}

	return user.Token, nil
}

func (p *botPersistence) FindUsername(username string) (*model.UserBot, error) {
	user := new(model.UserBot)
	coll := mgm.Coll(user)
	if err := coll.First(bson.M{"username": username}, user); err != nil {
		return nil, err
	}

	return user, nil
}

func BotInit(db *mongo.Database) bot.Persistence {
	return &botPersistence{db: db}
}
