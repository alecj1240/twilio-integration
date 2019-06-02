package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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

func CreateNumber(APIKeySID string, APIKeySecret string, AccountSID string, areaCode string) {
	apiURL := "https://" + APIKeySID + ":" + APIKeySecret + "@api.twilio.com/2010-04-01/Accounts/" + AccountSID + "/IncomingPhoneNumbers.json"
	data := url.Values{}
	data.Set("AreaCode", areaCode)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
}
