package handler

import (
	"encoding/json"
	"fmt"
	"log"
)

// GetUserNumbers - finds if the user has any existing phone numbers
func GetUserNumbers(APIKeySID string, APIKeySecret string, AccountSID string) []string {

	url := "https://" + APIKeySID + ":" + APIKeySecret + "@api.twilio.com/2010-04-01/Accounts/" + AccountSID + "/IncomingPhoneNumbers.json"
	getNumbers := GetRequest(url)

	var twilioNumbers TwilioNumberRequest
	err := json.Unmarshal(getNumbers, &twilioNumbers)
	if err != nil {
		fmt.Println(err)
	}

	userNumbers := make([]string, 0)
	for _, number := range twilioNumbers.IncomingPhoneNumbers {
		log.Println(number.PhoneNumber)
		userNumbers = append(userNumbers, number.PhoneNumber)
	}

	return userNumbers
}
