package repository

import (
	"gorm.io/gorm"
	"WEB1/internal/domain"
)

type UserModel struct {
	Name          string
	Email         string	`gorm:"unique"`
	HashedPassword string
	gorm.Model
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// Helper to convert domain.User to UserModel
func toUserModel(d domain.User) UserModel {
	return UserModel{
		Name:          d.Name,
		Email:         d.Email,
		HashedPassword: d.HashedPassword,
	}
}

// Helper to convert UserModel to domain.User
func toDomainUser(m UserModel) domain.User {
	return domain.User{
		ID:            m.ID,
		Name:          m.Name,
		Email:         m.Email,
		HashedPassword: m.HashedPassword,
		CreatedAt:     m.CreatedAt,
	}
}


func (r *userRepository) Create(user domain.User) (domain.User, error) {
	userModel := toUserModel(user)
	if err := r.db.Create(&userModel).Error; err != nil {
		return domain.User{}, err
	}
	return toDomainUser(userModel), nil
}

func (r *userRepository) Update(user domain.User) (domain.User, error) {
	userModel := toUserModel(user)
	if err := r.db.Save(&userModel).Error; err != nil {
		return domain.User{}, err
	}
	return toDomainUser(userModel), nil
}

func (r *userRepository) Delete(id uint) error {
	if err := r.db.Delete(&UserModel{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByID(id uint) (domain.User, error) {
	var userModel UserModel
	if err := r.db.First(&userModel, id).Error; err != nil {
		return domain.User{}, err
	}
	return toDomainUser(userModel), nil
}

func (r *userRepository) List() ([]domain.User, error) {
	var userModels []UserModel
	if err := r.db.Find(&userModels).Error; err != nil {
		return nil, err
	}

	var users []domain.User
	for _, userModel := range userModels {
		users = append(users, toDomainUser(userModel))
	}
	return users, nil
}