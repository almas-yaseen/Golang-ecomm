package handlers

import (
	"errors"
	"fmt"
	"ginapp/models"
	response "ginapp/reponse"
	"ginapp/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Signup(c *gin.Context) {

	var usersign models.SignupDetail

	if err := c.ShouldBindJSON(&usersign); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong formattttt 🙌", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	// CHEKING THE DATA ARE SENDED IN CORRECT FORMET OR NOT

	if err := validator.New().Struct(usersign); err != nil {

		errres := response.ClientResponse(404, "They are not in format", nil, err.Error())
		c.JSON(http.StatusBadGateway, errres)
		return
	}

	usercreate, err := usecase.UserSignup(usersign)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "user signup format error ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "User sign up succsed", usercreate, nil)
	c.JSON(http.StatusCreated, successRes)

}

func UserLoginWithPassword(c *gin.Context) {

	var LoginUser models.UserLogin

	if err := c.ShouldBindJSON(&LoginUser); err != nil {
		erres := response.ClientResponse(http.StatusBadGateway, "Login field provided in wrong way ", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return

	}

	////////

	if err := validator.New().Struct(LoginUser); err != nil {
		erres := response.ClientResponse(http.StatusBadGateway, "Login field was wrong formate ahn", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return
	}

	LogedUser, err := usecase.UserLogged(LoginUser)
	if errors.Is(err, models.ErrEmailNotFound) {

		erres := response.ClientResponse(http.StatusBadRequest, "invalid email", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return
	}

	if err != nil {

		erres := response.ClientResponse(500, "server error from usecase", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return
	}

	successres := response.ClientResponse(http.StatusCreated, "succesed login user", LogedUser, nil)

	c.JSON(http.StatusOK, successres)
}

func AddAddress(c *gin.Context) {

	fmt.Println("Hey from")

	user_id, _ := c.Get(models.User_id)

	var address models.AddressInfo
	if err := c.ShouldBindJSON(&address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	err := validator.New().Struct(address)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Constraints does not match", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	if err := usecase.Addaddress(user_id.(int), address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error from adding address", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return

	}

	successRes := response.SuccessClientResponse(http.StatusOK, "added address successfully")
	c.JSON(http.StatusOK, successRes)
}
