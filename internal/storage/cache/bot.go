package cache

import (
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"go-tellme/internal/constants/model"
	"go-tellme/internal/module/bot"
	"time"
)

const (
	NoExpiration time.Duration = -1
)

var (
	format      = "%s:%s"
	formatToken = "%s:token"
)

type botCaching struct {
	cache *memcache.Client
}

var userFormat = "client:%s"

func (u *botCaching) SaveID(ID string, user *model.UserBot) error {
	bob, err := json.Marshal(user)
	if err != nil {
		return err
	}
	item := &memcache.Item{
		Key:   fmt.Sprintf(userFormat, ID),
		Value: bob,
	}

	if err := u.cache.Set(item); err != nil {
		return err
	}

	return nil
}

func InitCache(cache *memcache.Client) bot.Caching {
	return &botCaching{cache}
}
