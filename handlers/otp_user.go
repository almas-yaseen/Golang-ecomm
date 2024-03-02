package handlers

import (
	"fmt"
	"ginapp/models"
	response "ginapp/reponse"
	"ginapp/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOtp(c *gin.Context) {

	var phone models.OTPData

	if err := c.ShouldBindJSON(&phone); err != nil {
		erres := response.ClientResponse(http.StatusBadGateway, "otp field provide wrong way", nil, err.Error())
		c.JSON(http.StatusBadRequest, erres)
		return
	}
	err := usecase.SendOtp(phone.PhoneNumber)

	if err != nil {

		erres := response.ClientResponse(http.StatusBadGateway, "cant send the otp", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "otp send successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func VerifyOTP(c *gin.Context) {

	fmt.Println("jey skjadnkasndk")

	var code models.VerifyData

	if err := c.ShouldBindJSON(&code); err != nil {
		erres := response.ClientResponse(http.StatusBadGateway, "json format is not correct", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return
	}
	users, err := usecase.VerifyOTP(code)
	if err != nil {
		erres := response.ClientResponse(http.StatusBadGateway, "could not  verify the otp", nil, err)
		c.JSON(http.StatusBadGateway, erres)
		return

	}
	successRes := response.ClientResponse(http.StatusOK, "Otp Verification Done", users, nil)
	c.JSON(http.StatusOK, successRes)

}
