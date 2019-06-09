package handler

import (
	"context"
	"fmt"
	"hackathon/mymongo"
	"hackathon/views"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Handler - handles the functions
func Handler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		respond(res, "")
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
			if len(userNumbers) > 0 {
				page := views.SelectNumber(userNumbers)
				respond(res, page)
				client.Disconnect(context.TODO())
				return
			} else {
				// if there are no numbers - show them the form to create a phonenumber
				respond(res, views.CreateNumber())
				client.Disconnect(context.TODO())
				return
			}
		}

		respond(res, "<Page><Container><Link href='https://zeit.co'>Visit Twilio Console for Account Sid</Link><BR /><Link href='https://www.twilio.com/console/runtime/api-keys/create'>Visit Twilio API Console to Create an API Key</Link><Input label='Twilio API Key SID' name='APIKeySID' /><Input label='Twilio API Key Secret' name='apiKeySecret' /><Input label='Twilio Account SID' name='TwilioAccountSID' /></Container><Container><Button action='createUser'>Submit</Button></Container></Page>")
		client.Disconnect(context.TODO())
		return

	case "createInstead":
		respond(res, views.CreateNumber())
		client.Disconnect(context.TODO())
		return
	case "createUser":
		// move onto index view after this is completed

		mymongo.CreateUser(client, zeitPayload.User.Id, zeitPayload.ClientState.APIKeySID, zeitPayload.ClientState.APIKeySecret, zeitPayload.ClientState.TwilioAccountSID)
		user := mymongo.FindUser(client, zeitPayload.User.Id)
		userNumbers := GetUserNumbers(user.TwilioKeySID, user.TwilioKeySecret, user.TwilioAccountSID)

		// if there are some numbers - ask if they want to use one of these
		if len(userNumbers) > 0 {
			page := views.SelectNumber(userNumbers)
			respond(res, page)
			client.Disconnect(context.TODO())
			return
		} else {
			// if there are no numbers - show them the form to create a phonenumber
			respond(res, views.CreateNumber())
			client.Disconnect(context.TODO())
			return
		}

	case "selectNumber":
		user := mymongo.FindUser(client, zeitPayload.User.Id)
		switch zeitPayload.ClientState.Language {
		case "Node.JS":
			respond(res, NodeJS(user.TwilioAccountSID, zeitPayload.ClientState.TwilioPhoneNumbers, user.TwilioKeySID, user.TwilioKeySecret))
		case "Php":
			respond(res, Php(user.TwilioAccountSID, zeitPayload.ClientState.TwilioPhoneNumbers, user.TwilioKeySID, user.TwilioKeySecret))
		case "Python":
			respond(res, Python(user.TwilioAccountSID, zeitPayload.ClientState.TwilioPhoneNumbers, user.TwilioKeySID, user.TwilioKeySecret))
		default:
			respond(res, NodeJS(user.TwilioAccountSID, zeitPayload.ClientState.TwilioPhoneNumbers, user.TwilioKeySID, user.TwilioKeySecret))
		}
		client.Disconnect(context.TODO())
		return

	case "createNumber":

		user := mymongo.FindUser(client, zeitPayload.User.Id)
		CreateNumber(user.TwilioKeySID, user.TwilioKeySecret, user.TwilioAccountSID, zeitPayload.ClientState.TwilioAreaCode)
		userNumbers := GetUserNumbers(user.TwilioKeySID, user.TwilioKeySecret, user.TwilioAccountSID)

		respond(res, views.SelectNumber(userNumbers))
		client.Disconnect(context.TODO())
		return

	default:
		// just send out a default view
		respond(res, "<Page><P>There is a default view</P></Page>")
		client.Disconnect(context.TODO())
		return
	}
}

func respond(res http.ResponseWriter, obj string) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization, Accept, Content-Type")
	res.Header().Set("Content-Type", "application/json")

	res.Write([]byte(obj))
}
