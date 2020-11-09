package repository

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"go-tellme/internal/module/bot"
	"go-tellme/platform/grpc/gen"
	"google.golang.org/grpc"
	"time"
)

type botRepository struct {
	persistence bot.Persistence
	amqp        *amqp.Connection
	grpc        *grpc.ClientConn
}

const (
	TimeoutRPC = 10 * time.Second
)

func (b *botRepository) IsUser(token string) error {
	if token != "abc123" {
		return errors.New("token not valid")
	}

	return nil
}

func (b *botRepository) ChatBot(payload *gen.ChatPayload) (*gen.ChatResponse, error) {
	log := logrus.WithFields(logrus.Fields{
		"domain":    "telegram",
		"action":    "RetrieveMessage",
		"rpcClient": "ChatBot",
	})

	ctx, cancel := context.WithTimeout(context.Background(), TimeoutRPC)
	defer cancel()

	service := gen.NewChatBotClient(b.grpc)
	response, err := service.RetrieveMessage(ctx, payload)

	if err != nil {
		log.WithField("type", "gRPC RetrieveMessage").Errorln(err)
		return nil, err
	}

	clientResponse := new(gen.ChatResponse)
	clientResponse.Message = response.Message
	clientResponse.Label = response.Label
	clientResponse.Accuracy = response.Accuracy

	return clientResponse, nil
}

func BotInit(persistence bot.Persistence, conn *amqp.Connection, clientConn *grpc.ClientConn) bot.Repository {
	return &botRepository{persistence: persistence, amqp: conn, grpc: clientConn}
}
