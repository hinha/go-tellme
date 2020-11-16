package repository

import "go-tellme/internal/module/client"

type clientRepository struct {
	persistence client.Persistence
}

func WebInit(persistence client.Persistence) client.Repository {
	return &clientRepository{persistence: persistence}
}
