package repository

import (
	"context"

	"github.com/ernestomr87/go-rest-websockets/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	Close() error
}

var implementations UserRepository

func SetRepository(repository UserRepository) {
	implementations = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementations.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementations.GetUserById(ctx, id)
}

func Close() error {
	return implementations.Close()
}
