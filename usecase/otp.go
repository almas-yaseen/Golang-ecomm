package usecase

import (
	"errors"
	"fmt"
	"ginapp/config"
	"ginapp/helper"
	"ginapp/models"
	"ginapp/repository"

	"github.com/jinzhu/copier"
)

func SendOtp(phone string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("cfg config")
		return errors.New("can not generate")

	}
	user, err := repository.FindUserByMobileNumber(phone)

	if err != nil {
		return errors.New("Error with server find user by mobile number")
	}

	if user == nil {
		return errors.New("Phone number not exist")

	}
	helper.TwilioSetUp(cfg.ACCOUNTSID, cfg.AUTHTOKEN)
	_, err = helper.TwilioSendOtp(phone, cfg.SERVICESSID)

	if err != nil {
		fmt.Println("here is my error ", cfg.SERVICESSID, "my token", cfg.AUTHTOKEN, "my account sid:=", cfg.ACCOUNTSID)
		return errors.New("error occured while generating otp co")
	}
	return nil

}

func VerifyOTP(code models.VerifyData) (models.TokenUser, error) {
	cfg, err := config.LoadConfig()

	if err != nil {
		return models.TokenUser{}, err
	}
	helper.TwilioSetUp(cfg.ACCOUNTSID, cfg.AUTHTOKEN)
	err = helper.TwilioVerifyOTP(cfg.SERVICESSID, code.Code, code.User.PhoneNumber)
	if err != nil {
		return models.TokenUser{}, err
	}
	userDetails, err := repository.UserDetailsUsingPhone(code.User.PhoneNumber)
	if err != nil {
		return models.TokenUser{}, err
	}
	accessToken, err := helper.GenerateAccessToken(userDetails)
	if err != nil {
		return models.TokenUser{}, err
	}
	refreshToken, err := helper.GenerateRefreshToken(userDetails)
	if err != nil {
		return models.TokenUser{}, err
	}

	var user models.SignupDetailResponse
	err = copier.Copy(&user, &userDetails)
	if err != nil {
		return models.TokenUser{}, err
	}
	return models.TokenUser{
		Users:        user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
