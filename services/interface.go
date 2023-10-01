package services

import (
	"base-gin-go/models"
	"context"
)

type UserServicer interface {
	GetUser(ctx context.Context, id int) (user *models.User, err error)
}
