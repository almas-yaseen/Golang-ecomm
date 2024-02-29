package usecase

import (
	"errors"
	"fmt"
	"ginapp/helper"
	"ginapp/models"
	"ginapp/repository"
	"net/mail"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(user models.SignupDetail) (*models.TokenUser, error) {

	email, err := repository.CheckingEmailValidation(user.Email)

	if err != nil {
		fmt.Println(err)
		return &models.TokenUser{}, errors.New("error with the singup server")

	}

	if email != nil {
		return &models.TokenUser{}, errors.New("email is already exisit ")
	}

	phone, err := repository.CheckingPhoneExists(user.Phone)

	if err != nil {

		return &models.TokenUser{}, errors.New("p_server have issue ")

	}
	if phone != nil {
		return &models.TokenUser{}, errors.New("phone number is already exist ")

	}

	//   Passoword Hash

	hashedPassword, err := helper.PasswordHashing(user.Password)

	if err != nil {
		return &models.TokenUser{}, errors.New("hash_server have issue")
	}

	user.Password = hashedPassword

	dataInsert, err := repository.SignupInsert(user)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not add User ")
	}

	fmt.Println("data inserted in signup :", dataInsert)

	// CREATING A JWT TOKEN FOR THE NEW USER\\

	accessToken, err := helper.GenerateAccessToken(dataInsert)

	if err != nil {
		fmt.Println("kljasndkaskdkasdkbasdkhbaksdbkhabsdkjbhsd", err)
		return &models.TokenUser{}, errors.New("can't create a acces token")
	}

	refershToken, err := helper.GenerateRefreshToken(dataInsert)

	if err != nil {
		return &models.TokenUser{}, errors.New("can't create a Refersh token")

	}

	return &models.TokenUser{
		Users:        dataInsert,
		AccessToken:  accessToken,
		RefreshToken: refershToken,
	}, nil

}

func UserLogged(user models.UserLogin) (*models.TokenUser, error) {

	_, err := mail.ParseAddress(user.Email)
	if err != nil {

		return &models.TokenUser{}, errors.New("EMAIL SHOULD BE CORRECT FORMAT ")

	}

	email, err := repository.CheckingEmailValidation(user.Email)

	if err != nil {
		return &models.TokenUser{}, errors.New("SERVER ERROR  from : checking-email-validation")

	}

	if email == nil {
		return &models.TokenUser{}, models.ErrEmailNotFound

	}

	userDetails, err := repository.FindUserDetailByEmail(user)
	if err != nil {
		return &models.TokenUser{}, err

	}

	// CHECKING THE HASSED PASSWORD

	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password))
	if err != nil {
		fmt.Println("HEyyyy")
		return &models.TokenUser{}, errors.New("hassed password not matching")
	}

	var user_details models.SignupDetailResponse

	err = copier.Copy(&user_details, &userDetails)
	if err != nil {
		return &models.TokenUser{}, errors.New("error in the copier")
	}

	//TOKEN....

	accessToken, err := helper.GenerateAccessToken(user_details)

	if err != nil {
		return &models.TokenUser{}, errors.New("Could not create access token due to error")

	}
	refreshToken, err := helper.GenerateRefreshToken(user_details)

	if err != nil {
		return &models.TokenUser{}, errors.New("could not create refresh token due to error")
	}

	return &models.TokenUser{

		Users:        user_details,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
