package repository

import (
	"fmt"
	"go-echo-experiment/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByID(ID uint) (*model.User, error) {
	user := model.User{}
	err := r.DB.Where("id = ?", ID).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := r.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user model.User) (*model.User, error) {
	data := model.User{}

	err := r.DB.Create(&user).Error

	if err != nil {
		fmt.Printf("Error1: %v\n", err)
		return nil, err
	}

	id, errId := r.GetUserByID(user.ID)
	if errId != nil {
		fmt.Printf("Error2: %v\n", errId)
		return nil, errId
	}

	data = *id

	return &data, nil
}

func (r *UserRepository) LoginUser(user *model.User) (*model.User, error) {
	data := model.User{}
	err := r.DB.
		Select("id, name, email, password, salt").
		Where("email = ?", &user.Email).
		First(&data).Error

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return &data, nil
}
