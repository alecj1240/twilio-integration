package handler

import (
	"context"
	"fmt"
	"hackathon/mymongo"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Handler - handles the functions
func Handler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		Respond(res, nil)
		return
	}

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://alecjones:C3ITZEfRfY4R%404F808H@zeit-integration-0g5df.mongodb.net/test")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	zeitPayload := FormatRequest(req)
	fmt.Println(zeitPayload)

	// switch statement for what the user is doing in the integration
	switch zeitPayload.Action {
	case "view":
		user := mymongo.FindUser(client, zeitPayload.User.Id)
		if user.ZeitID != "" {
			// Get the numbers from Twilio
			userNumbers := GetUserNumbers(user.TwilioKeySID, user.TwilioKeySecret, user.TwilioAccountSID)

			// if there are some numbers - ask if they want to use one of these
			if len(userNumbers) > 1 {
				page := "<Page><H1>Select A Twilio Number</H1><Select name='twilioPhoneNumbers' value='selectedValue' action='selectNumber'>"
				for _, number := range userNumbers {
					page = page + "<Option name='twilioPhoneNumbers' value='" + number + "' caption='" + number + "' />"
				}
				page = page + "</Select><Button action='selectNumber'>Use This Number</Button></Page>"

				Respond(res, page)
				client.Disconnect(context.TODO())
				return

			} else {
				// if there are no numbers - show them the form to create a phonenumber
				Respond(res, "<Page><Container><H1>Create A Twilio Number</H1><Input label='Area Code of Your Desired PhoneNumber' name='areaCode' /></Container><Container><P>Note: you may be subject to charges from Twilio</P><Button action='createNumber'>Create Number</Button></Container></Page>")
				client.Disconnect(context.TODO())
				return
			}

			Respond(res, "<Page><P>Code Will Show Here</P></Page>")
		}

		Respond(res, "<Page><Container><Link href='https://zeit.co'>Visit Twilio Console for Account Sid</Link><BR /><Link href='https://www.twilio.com/console/runtime/api-keys/create'>Visit Twilio API Console to Create an API Key</Link><Input label='Twilio API Key SID' name='APIKeySID' /><Input label='Twilio API Key Secret' name='apiKeySecret' /><Input label='Twilio Account SID' name='TwilioAccountSID' /></Container><Container><Button action='createUser'>Submit</Button></Container></Page>")
		client.Disconnect(context.TODO())
		return

	case "createUser":
		// move onto index view after this is completed
		mymongo.CreateUser(client, zeitPayload.User.Id, zeitPayload.ClientState.APIKeySID, zeitPayload.ClientState.APIKeySecret, zeitPayload.ClientState.TwilioAccountSID)
		user := mymongo.FindUser(client, zeitPayload.User.Id)
		// Get the numbers from Twilio
		userNumbers := GetUserNumbers(user.TwilioKeySID, user.TwilioKeySecret, user.TwilioAccountSID)

		// if there are some numbers - ask if they want to use one of these
		if len(userNumbers) > 0 {
			page := "<Page><H1>Select A Twilio Number</H1><Select name='twilioPhoneNumber' action='selectNumber'>"
			for _, number := range userNumbers {

				page = page + "<Option name='twilioPhoneNumber' value='" + number + "' caption='" + number + "' />"
			}
			page = page + "</Select><Button action='selectNumber'>Use This Number</Button></Page>"

			Respond(res, page)
			client.Disconnect(context.TODO())
			return
		} else {
			// if there are no numbers - show them the form to create a phonenumber
			Respond(res, "<Page><Container><H1>Create A Twilio Number</H1><Input label='Area Code of Phone Number (3 digits)' name='areaCode' /></Container><Container><P>Note: you may be subject to charges from Twilio</P><Button action='createNumber'>Create Number</Button></Container></Page>")
			client.Disconnect(context.TODO())
			return
		}

	case "selectNumber":
		Respond(res, "<Page><P>Code Will Show Here</P></Page>")
		return
	case "createNumber":
		// TODO: Create the Twilio phone number
		user := mymongo.FindUser(client, zeitPayload.User.Id)

		CreateNumber(user.TwilioKeySID, user.TwilioKeySecret, user.TwilioAccountSID, zeitPayload.ClientState.TwilioAreaCode)
		Respond(res, "<Page><P>Code Will Show Here</P></Page>")
		client.Disconnect(context.TODO())
		return
	default:
		// just send out a default view
		Respond(res, "<Page><P>There is a default view</P></Page>")
		client.Disconnect(context.TODO())
		return
	}
}
