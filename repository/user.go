package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var userLogin model.User
	err := r.db.Table("users").Find(&userLogin, "email = ?", email).Error
	if err != nil {
		return model.User{}, err
	}
	return userLogin, nil 
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	var userCategories []model.UserTaskCategory
	err := r.db.Table("users").
		Select("users.id, users.fullname, users.email, tasks.title as task, tasks.deadline, tasks.priority, tasks.status, categories.name as category").
		Joins("JOIN tasks ON tasks.user_id = users.id").
		Joins("JOIN categories ON categories.id = tasks.category_id").
		Scan(&userCategories).Error
	if err != nil {
		return nil, err
	}
	return userCategories, nil 
}
