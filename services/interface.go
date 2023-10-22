package services

import (
	"base-gin-go/models"
	"context"
)

type UserServicer interface {
	Login(ctx context.Context, req models.RequestLogin) (string, error)
	GetUser(ctx context.Context, id int) (*models.User, error)
}
