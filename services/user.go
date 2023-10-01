package services

import (
	"base-gin-go/models"
	"base-gin-go/repositories"
	"context"
)

type UserService interface {
	// Login(ctx context.Context, req models.RequestLogin) (string, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUser(ctx context.Context, id int) (*models.User, error)
	CreateUser(ctx context.Context, u models.User) error
}

// func (s *AppService) Login(ctx context.Context, req models.RequestLogin) (token string, err error) {
// 	user, err := re.UserRepo.GetUserByEmail(req.Email)
// 	if err != nil {
// 		return "", fmt.Errorf("login fail")
// 	}
// 	result := hash.CompareHashAndPassword(user.Password, req.Password)
// 	if !result {
// 		return "", fmt.Errorf("login fail")
// 	}
// 	return jwt.GenerateToken(user.ID)
// }

// func (s *AppService) GetUsers(ctx context.Context) (users []models.User, err error) {
// 	users, err = repositories.GetUsers(ctx, s.db, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

func (s *AppService) GetUser(ctx context.Context, id int) (user *models.User, err error) {
	user, err = repositories.GetUser(ctx, s.db, id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

// func (s *AppService) CreateUser(ctx context.Context, u models.User) (err error) {
// 	u.HashPassword(u.Password)
// 	_, err = uc.UserRepo.CreateUser(u)
// 	if err != nil {
// 		return err
// 	}

// 	return
// }
