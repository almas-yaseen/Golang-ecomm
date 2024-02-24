package repository

import (
	"errors"
	"fmt"
	"ginapp/database"
	"ginapp/domain"

	"github.com/jinzhu/gorm"
)

func CheckEmailValidation(email string) (*domain.User, error) {

	var user domain.User
	result := database.DB.Where(&domain.User{Email: email}).First(&user)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error

	}
	fmt.Println("User details was ", user)
	return &user, nil

}

func CheckingPhoneExists(phone string) (*domain.User, error) {

	var user domain.User
	result := database.DB.Where(&domain.User{Phone: phone}).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func Signup
