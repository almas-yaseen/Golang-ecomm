package helper

import (
	"errors"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient

func TwilioSetUp(username string, password string) {

	client = twilio.NewRestClientWithParams(twilio.ClientParams{

		Username: username,
		Password: password,
	})

}

func TwilioSendOtp(phone string, serviceID string) (string, error) {
	params := &twilioApi.CreateVerificationParams{} //create  new verfication params from the twilio api
	params.SetTo("+91" + phone)
	params.SetChannel("sms")

	fmt.Println("service id: and params:=", serviceID, params)

	// verify v2 service provided by twilio that allows to send verification
	// --create verfication-> this is the method provided by verfiy v2 service it is used to create new verfication request

	resp, err := client.VerifyV2.CreateVerification(serviceID, params)

	if err != nil {
		fmt.Println("hey from send otp", err)

		return "", err
	}
	return *resp.Sid, nil
}
func TwilioVerifyOTP(serviceID string, code string, phone string) error {

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceID, params)
	fmt.Println("resp status", resp)
	if err != nil {
		return err
	}
	if *resp.Status == "approved" {
		return nil
	}
	return errors.New("Failed to validate otp")
}
