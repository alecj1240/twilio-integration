package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"hackathon/mongo"
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
	AccountSID string `json:"accountSID"`
	AuthToken  string `json:"authToken"`
}

type UserRequest struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Handler - handles the functions
func Handler(res http.ResponseWriter, req *http.Request) {

	if req.Method == "OPTIONS" {
		respond(res, nil)
		return
	}
	msg := FormatRequest(req)
	fmt.Println(msg)

	mongo.CreateUser("hello", "there")

	respond(res, "<Page><Container><Input label='Twilio Account SID' name='accountSID' /><Input label='Twilio Auth Token' name='authToken' /></Container><Container><Button action='data'>Submit</Button></Container></Page>")
}

// Respond - used to make it easy to respond to requests
func respond(res http.ResponseWriter, obj interface{}) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization, Accept, Content-Type")
	res.Header().Set("Content-Type", "application/json")

	j := json.NewEncoder(res)
	j.SetEscapeHTML(false)
	j.Encode(obj)
	res.Write([]byte("\n"))
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
