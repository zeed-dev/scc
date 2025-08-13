package services

import (
	"errors"
	"smart-command-center-backend/config"
	"smart-command-center-backend/models"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
	RoleID   uint `json:"role_id"`
}

type UpdateUserInput struct {
	Name     string
	Email    string
	Password string
	RoleID   uint `json:"role_id"`
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Preload("Role").Find(&users).Error
	return users, err
}

func GetUserByID(id string) (models.User, error) {
	var user models.User
	err := config.DB.Preload("Role").First(&user, id).Error
	return user, err
}

func CreateUser(input CreateUserInput) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	var exists bool
	err = config.DB.Model(&models.Role{}).
		Select("count(*) > 0").
		Where("id = ?", input.RoleID).
		Find(&exists).Error
	if err != nil {
		return models.User{}, err
	}
	if !exists {
		return models.User{}, errors.New("role not found")
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		RoleID:   input.RoleID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	if err := config.DB.Preload("Role").First(&user, user.ID).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdateUser(id string, input UpdateUserInput) (models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.RoleID != 0 {
		var exists bool
		err := config.DB.Model(&models.Role{}).
			Select("count(*) > 0").
			Where("id = ?", input.RoleID).
			Find(&exists).Error
		if err != nil {
			return models.User{}, err
		}
		if !exists {
			return models.User{}, errors.New("role not found")
		}
		user.RoleID = input.RoleID
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, err
		}
		user.Password = string(hashedPassword)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	if err := config.DB.Preload("Role").First(&user, user.ID).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func DeleteUser(id string) error {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
