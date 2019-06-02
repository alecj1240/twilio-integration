package handler

import (
	"fmt"
	"hackathon/mongo"
	"log"
	"net/http"
)

// Handler - handles the functions
func Handler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		Respond(res, nil)
		return
	}

	prettyReq := PrettyRequest(req)
	log.Println(prettyReq)
	zeitPayload := FormatRequest(req)
	fmt.Println(zeitPayload)

	// switch statement for what the user is doing in the integration
	switch zeitPayload.Action {
	case "view":
		user := mongo.FindUser(zeitPayload.User.Id)
		if user.ZeitID != "" {
			// Get the numbers from Twilio
			userNumbers := GetUserNumbers(user.TwilioKeySID, user.TwilioKeySecret, user.TwilioAccountSID)

			// if there are some numbers - ask if they want to use one of these
			if len(userNumbers) > 0 {
				page := "<Page><H1>Select A Twilio Number</H1><Select name='twilioPhoneNumbers' value='selectedValue' action='selectNumber'>"
				for _, number := range userNumbers {

					page = page + "<Option name='twilioPhoneNumbers' value='" + number + "' caption='" + number + "' />"
				}
				page = page + "</Select><Button action='selectNumber'>Use This Number</Button></Page>"

				Respond(res, page)
				return
			}

			// if there are no numbers - show them the form to create a phonenumber
			Respond(res, "<Page><Container><H1>Create A Twilio Number</H1><Input label='Area Code of Your Desired PhoneNumber' name='areaCode' /></Container><Container><P>Note: you may be subject to charges from Twilio</P><Button action='createNumber'>Create Number</Button></Container></Page>")
			return
		}

		Respond(res, "<Page><Container><Input label='Twilio API Key SID' name='APIKeySID' /><Input label='Twilio API Key Secret' name='apiKeySecret' /><Input label='Twilio Account SID' name='TwilioAccountSID' /></Container><Container><Button action='createUser'>Submit</Button></Container></Page>")
		return

	case "createUser":
		// move onto index view after this is completed
		mongo.CreateUser(zeitPayload.User.Id, zeitPayload.ClientState.APIKeySID, zeitPayload.ClientState.APIKeySecret, zeitPayload.ClientState.TwilioAccountSID)
		Respond(res, "<Page><P>There is a user</P></Page>")
		return

	case "selectNumber":
		Respond(res, "<Page><P>Code Will Show Here</P></Page>")
	case "createNumber":

	default:
		// just send out a default view
		Respond(res, "<Page><P>There is a default view</P></Page>")
		return
	}
}
