package repositories

import (
	"base-gin-go/models"
	"context"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers() ([]models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUser(id int) (*models.User, error)
	CreateUser(u models.User) (*models.User, error)
}

// func GetUserByEmail(email string) (user models.User, err error) {
// 	result := r.DB.Where("email = ?", email).First(&user)
// 	return user, result.Error
// }

func GetUser(ctx context.Context, db *gorm.DB, id int) (user *models.User, err error) {
	if err = db.First(&user, id).Error; err != nil {
		return nil, err
	}
	user.Password = ""
	return
}

// func CreateUser(u models.User) (user *models.User, err error) {
// 	user = &models.User{
// 		Username: u.Username,
// 		Email:    u.Email,
// 		Password: u.Password,
// 	}
// 	result := r.DB.Create(user)

// 	if result.Error != nil {
// 		err = fmt.Errorf("[repo.User.CreateUser] failed: %w", result.Error)
// 		return nil, err
// 	}

// 	return
// }
