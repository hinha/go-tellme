package cache

import (
	"errors"
	"fmt"
	gcache "github.com/patrickmn/go-cache"
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
	cache *gcache.Cache
}

func (b *botCaching) SaveToken(username, token string) error {
	key := fmt.Sprintf(formatToken, username)
	b.cache.Set(key, token, NoExpiration)
	return nil
}

func (b *botCaching) SaveAction(username, action string) error {
	key := fmt.Sprintf(format, username, action)
	b.cache.Set(key, "ok", NoExpiration)
	return nil
}

func (b *botCaching) GetToken(username string) (string, error) {
	var token string
	x, found := b.cache.Get(formatToken)
	if !found {
		return "", errors.New("token not found")
	}
	token = x.(string)

	return token, nil
}

func (b *botCaching) GetAction(username, action string) error {
	key := fmt.Sprintf(format, username, action)
	_, got := b.cache.Get(key)
	if !got {
		return errors.New("action not found")
	}
	return nil
}

func (b *botCaching) GetStartAction(username string) error {
	format := fmt.Sprintf("%s:/start", username)
	b.cache.Get(format)

	var foo string
	if x, found := b.cache.Get(format); found {
		foo = x.(string)
	}

	if foo != "ok" {
		return errors.New("not permissions")
	}

	return nil
}

func InitCache(cache *gcache.Cache) bot.Caching {
	return &botCaching{cache}
}
