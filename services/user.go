package services

import (
	"base-gin-go/models"
	"base-gin-go/pkg/hash"
	"base-gin-go/pkg/jwt"
	"base-gin-go/repositories"
	"context"
	"fmt"
)

func (s *AppService) Login(ctx context.Context, req models.RequestLogin) (token string, err error) {
	user, err := repositories.GetByEmail(ctx, s.db, req.Email)
	if err != nil {
		return "", fmt.Errorf("login fail")
	}
	result := hash.CompareHashAndPassword(user.Password, req.Password)
	if !result {
		return "", fmt.Errorf("login fail")
	}
	return jwt.GenerateToken(user.ID)
}

func (s *AppService) GetUser(ctx context.Context, id int) (user *models.User, err error) {
	user, err = repositories.Detail(ctx, s.db, id)

	if err != nil {
		return nil, err
	}
	return user, nil
}
