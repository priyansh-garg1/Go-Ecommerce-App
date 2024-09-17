package repository

import (
	"errors"
	"go-ecommerce-app/internal/domain"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {
	err := r.db.Create(&usr).Error
	if err != nil {
		log.Printf("Create User error: %v", err)
		return domain.User{}, errors.New("failed to create user error")
	}
	return usr, nil
}
func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("Find User error: %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}
func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, id).Error
	if err != nil {
		log.Printf("Find User error: %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}
func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User
	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error
	if err != nil {
		log.Printf("Update User error: %v", err)
		return domain.User{}, errors.New("failed to update user error")
	}
	return user, nil
}
