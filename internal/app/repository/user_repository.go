package repository

import (
	"gorm.io/gorm"
	"WEB1/internal/domain"
)

type UserModel struct {
	ID            uint       `gorm:"primaryKey"`
	Name          string
	Email         string	`gorm:"unique"`
	HashedPassword string
	gorm.Model
}

type userRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepository{db: db}
}

func (r *userRepository) Create(user domain.User) (domain.User, error) {
    userModel := UserModel{
        Name:             user.Name,
        Email:            user.Email,
        HashedPassword:   user.HashedPassword,
    }
    if err := r.db.Create(&userModel).Error; err != nil {
        return domain.User{}, err
    }
    
   
    return domain.User{
        ID:             userModel.ID,
        Name:           userModel.Name,
        Email:          userModel.Email,
        HashedPassword: userModel.HashedPassword,
        CreatedAt:      userModel.CreatedAt,
    }, nil
}

func (r *userRepository) GetByID(id uint) (domain.User, error) {
	var user UserModel
	if err := r.db.First(&user, id).Error; err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		HashedPassword: user.HashedPassword,
		CreatedAt:     user.CreatedAt,
	}, nil
}

func (r *userRepository) List() ([]domain.User, error) {
	var users []UserModel
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	var result []domain.User
	for _, user := range users {
		result = append(result, domain.User{
			ID:            user.ID,
			Name:          user.Name,
			Email:         user.Email,
			HashedPassword: user.HashedPassword,
			CreatedAt:     user.CreatedAt,
		})
	}
	return result, nil
}