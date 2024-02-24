package usecase

import "ginapp/models"

func UserSignup(user models.SignupDetail)(*models.TokenUser,error)

email,err := respository.CheckingEmailValidation(user.Email)

if err != nil{


	return &models.TokenUser{},errors.New("error with signup server")
}

if email !=nil{

	return &models.TokenUser{},errors.New("email is already exist")

}

phone,err := repository.CheckPhoneExists(user.Phone)

if err!=nil{
	return &models.TokenUser{},errors.New("p_server has issue")
}

if phone !=nil{
	return &models.TokenUser{},errors.New("Phone number already exist")
}
hashedPassword,err := helper.PasswordHashing(user.Password)

if err!=nil{

	return &models.TokenUser{},errors.New("hash_server has issue")
}