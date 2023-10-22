package repositories

import (
	"base-gin-go/models"
	"context"

	"gorm.io/gorm"
)

type UserRepo interface {
	List() ([]models.User, error)
	GetByEmail(ctx context.Context, db *gorm.DB, email string) (models.User, error)
	Detail(ctx context.Context, db *gorm.DB, id int) (*models.User, error)
	CreateUser(u models.User) (*models.User, error)
}

func GetByEmail(ctx context.Context, db *gorm.DB, email string) (user models.User, err error) {
	result := db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func Detail(ctx context.Context, db *gorm.DB, id int) (user *models.User, err error) {
	if err = db.First(&user, id).Error; err != nil {
		return nil, err
	}
	user.Password = ""
	return
}
