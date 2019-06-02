package handler

import (
	"fmt"
	"net/http"
	"hackathon/mongo"
	
)



// Handler - handles the functions
func Handler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		Respond(res, nil)
		return
	}

	zeitPayload := FormatRequest(req)
	fmt.Println(zeitPayload)

	switch zeitPayload.Action {
		case "view":
			
			if mongo.FindUser(zeitPayload.User.Id) {
				Respond(res, "<Page><P>There is a user</P></Page>")
				return
			}

			Respond(res, "<Page><Container><Input label='Twilio Account SID' name='accountSID' /><Input label='Twilio API Key' name='apiKey' /></Container><Container><Button action='createUser'>Submit</Button></Container></Page>")
			return
		
		case "createUser":
			// move onto index view after this is completed
			mongo.CreateUser(zeitPayload.User.Id, zeitPayload.ClientState.AccountSID, zeitPayload.ClientState.APIKey)
			Respond(res, "<Page><P>There is a user</P></Page>")
			return

		default: 
			// just send out a default view
			Respond(res, "<Page><P>There is a default view</P></Page>")
			return
	}
}

