package repositories

import (
	"gorm.io/gorm"
	"github.com/KshitijBhardwaj18/Orbix/shared/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db:db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil,err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}
