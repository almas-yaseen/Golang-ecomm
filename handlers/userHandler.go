package handlers

import (
	"ginapp/models"
	response "ginapp/reponse"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Signup(c *gin.Context) {

	var usersign models.SignupDetail

	if err := c.ShouldBindJSON(&usersign); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := validator.New().Struct(usersign); err != nil {

		errres := response.ClientResponse(404, "They are not formatted", nil, err.Error()())
		c.JSON(http.StatusBadGateway, errres)
		return
	}

	usercreate, err := usecase.UserSignup(usersign)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user signup format error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "user signup success", usercreate, nil)
	c.JSON(http.StatusCreated, successRes)

}
