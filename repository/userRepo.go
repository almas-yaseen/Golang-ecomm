package repository

import (
	"errors"
	"fmt"
	"ginapp/database"
	"ginapp/domain"
	"ginapp/models"

	"gorm.io/gorm"
)

func CheckingEmailValidation(email string) (*domain.User, error) {

	var user domain.User

	result := database.DB.Where(&domain.User{Email: email}).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil

		}
		return nil, result.Error
	}
	fmt.Println("user details was :", user)
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

func SignupInsert(user models.SignupDetail) (models.SignupDetailResponse, error) {
	var signupres models.SignupDetailResponse
	err := database.DB.Raw("INSERT INTO users(firstname, lastname, email, phone, password) VALUES (?,?,?,?,?) RETURNING id,firstname,lastname,email,phone", user.FirstName, user.LastName, user.Email, user.Phone, user.Password).Scan(&signupres).Error
	if err != nil {
		fmt.Println(err, "dd")
		return models.SignupDetailResponse{}, err
	}
	fmt.Println("signup inserted data are", signupres)
	return signupres, nil

}
func FindUserDetailByEmail(user models.UserLogin) (models.UserLoginResponse, error) {

	var UserDetails models.UserLoginResponse

	err := database.DB.Raw(
		`select * from users where email = ? and blocked = false`, user.Email).Scan(&UserDetails).Error

	if err != nil {
		return models.UserLoginResponse{}, errors.New("got an error fron ! searching users by email")

	}

	return UserDetails, nil
}

func FindUserByMobileNumber(phone string) (*domain.User, error) {
	var User domain.User
	result := database.DB.Where(&domain.User{Phone: phone}).First(&User)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error

	}
	return &User, nil

}
