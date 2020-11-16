package persistence

import (
	"github.com/Kamva/mgm/v3"
	"go-tellme/internal/constants/model"
	"go-tellme/internal/module/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type clientPersistence struct {
	db *mongo.Database
}

func (c *clientPersistence) UniqueEmail(email string) error {
	user := new(model.User)
	coll := mgm.Coll(user)
	if err := coll.First(bson.M{"email": email}, user); err != nil {
		return err
	}

	return nil
}

func (c *clientPersistence) InsertAccount(user *model.User) error {
	coll := mgm.Coll(user)
	if err := coll.Create(user); err != nil {
		return err
	}

	return nil
}

func (c *clientPersistence) GetAccount(email string) (*model.User, error) {
	user := new(model.User)
	coll := mgm.Coll(user)
	if err := coll.First(bson.M{"email": email}, user); err != nil {
		return nil, err
	}

	return user, nil
}

func ClientInit(db *mongo.Database) client.Persistence {
	return &clientPersistence{db: db}
}
