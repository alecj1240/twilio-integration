package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Payload struct {
	Action          string      `json:"action"`
	ClientState     ClientState `json:"clientState"`
	Slug            string      `json:"configurationId"`
	ConfigurationId string      `json:"configurationId"`
	IntegrationId   string      `json:"integrationId"`
	Team            string      `json:"team"`
	User            UserRequest `json:"user"`
	Project         string      `json:"project"`
	Token           string      `json:"token"`
}

type ClientState struct {
	APIKeySID          string `json:"APIKeySID"`
	APIKeySecret       string `json:"APIKeySecret"`
	TwilioAccountSID   string `json:"TwilioAccountSID"`
	TwilioPhoneNumbers string `json:"twilioPhoneNumber"`
	TwilioAreaCode     string `json:"areaCode"`
}

type UserRequest struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type TwilioNumberRequest struct {
	FirstPageURI         string      `json:"first_page_uri"`
	End                  int         `json:"end"`
	PreviousPageURI      interface{} `json:"previous_page_uri"`
	IncomingPhoneNumbers []struct {
		Sid                 string      `json:"sid"`
		AccountSid          string      `json:"account_sid"`
		FriendlyName        string      `json:"friendly_name"`
		PhoneNumber         string      `json:"phone_number"`
		VoiceURL            string      `json:"voice_url"`
		VoiceMethod         string      `json:"voice_method"`
		VoiceFallbackURL    interface{} `json:"voice_fallback_url"`
		VoiceFallbackMethod string      `json:"voice_fallback_method"`
		VoiceCallerIDLookup bool        `json:"voice_caller_id_lookup"`
		DateCreated         string      `json:"date_created"`
		DateUpdated         string      `json:"date_updated"`
		SmsURL              string      `json:"sms_url"`
		SmsMethod           string      `json:"sms_method"`
		SmsFallbackURL      string      `json:"sms_fallback_url"`
		SmsFallbackMethod   string      `json:"sms_fallback_method"`
		AddressRequirements string      `json:"address_requirements"`
		Beta                bool        `json:"beta"`
		Capabilities        struct {
			Voice bool `json:"voice"`
			Sms   bool `json:"sms"`
			Mms   bool `json:"mms"`
			Fax   bool `json:"fax"`
		} `json:"capabilities"`
		VoiceReceiveMode     string      `json:"voice_receive_mode"`
		StatusCallback       string      `json:"status_callback"`
		StatusCallbackMethod string      `json:"status_callback_method"`
		APIVersion           string      `json:"api_version"`
		VoiceApplicationSid  interface{} `json:"voice_application_sid"`
		SmsApplicationSid    string      `json:"sms_application_sid"`
		Origin               string      `json:"origin"`
		TrunkSid             interface{} `json:"trunk_sid"`
		EmergencyStatus      string      `json:"emergency_status"`
		EmergencyAddressSid  interface{} `json:"emergency_address_sid"`
		AddressSid           interface{} `json:"address_sid"`
		IdentitySid          interface{} `json:"identity_sid"`
		URI                  string      `json:"uri"`
		Status               string      `json:"status"`
	} `json:"incoming_phone_numbers"`
	URI         string      `json:"uri"`
	PageSize    int         `json:"page_size"`
	Start       int         `json:"start"`
	NextPageURI interface{} `json:"next_page_uri"`
	Page        int         `json:"page"`
}

type NumberCreated struct {
	Sid                 string      `json:"sid"`
	AccountSid          string      `json:"account_sid"`
	FriendlyName        string      `json:"friendly_name"`
	PhoneNumber         string      `json:"phone_number"`
	VoiceURL            string      `json:"voice_url"`
	VoiceMethod         string      `json:"voice_method"`
	VoiceFallbackURL    interface{} `json:"voice_fallback_url"`
	VoiceFallbackMethod string      `json:"voice_fallback_method"`
	VoiceCallerIDLookup bool        `json:"voice_caller_id_lookup"`
	DateCreated         string      `json:"date_created"`
	DateUpdated         string      `json:"date_updated"`
	SmsURL              string      `json:"sms_url"`
	SmsMethod           string      `json:"sms_method"`
	SmsFallbackURL      string      `json:"sms_fallback_url"`
	SmsFallbackMethod   string      `json:"sms_fallback_method"`
	AddressRequirements string      `json:"address_requirements"`
	Beta                bool        `json:"beta"`
	Capabilities        struct {
		Voice bool `json:"voice"`
		Sms   bool `json:"sms"`
		Mms   bool `json:"mms"`
		Fax   bool `json:"fax"`
	} `json:"capabilities"`
	VoiceReceiveMode     string      `json:"voice_receive_mode"`
	StatusCallback       string      `json:"status_callback"`
	StatusCallbackMethod string      `json:"status_callback_method"`
	APIVersion           string      `json:"api_version"`
	VoiceApplicationSid  interface{} `json:"voice_application_sid"`
	SmsApplicationSid    string      `json:"sms_application_sid"`
	Origin               string      `json:"origin"`
	TrunkSid             interface{} `json:"trunk_sid"`
	EmergencyStatus      string      `json:"emergency_status"`
	EmergencyAddressSid  interface{} `json:"emergency_address_sid"`
	AddressSid           interface{} `json:"address_sid"`
	IdentitySid          interface{} `json:"identity_sid"`
	URI                  string      `json:"uri"`
	Status               string      `json:"status"`
}

// FormatRequest generates ascii representation of a request
func FormatRequest(req *http.Request) Payload {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	var msg Payload
	err = json.Unmarshal(b, &msg)
	if err != nil {
		fmt.Println(err)
	}

	return msg
}

func PostRequest(url string) []byte {
	vard := map[string]string{}
	requestBody, err := json.Marshal(vard)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

// GetRequest - HTTP GET functionality
func GetRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
