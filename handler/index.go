package handler

import (
	"context"
	"encoding/json"
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
		respond(res, nil)
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

				page := "<Page><Container><H1>Select A Phone Number</H1><BR /><P><B>please type in one of the phone numbers listed below</B></P>"
				for _, number := range userNumbers {
					page = page + "<P>" + number + "</P><BR />"
				}
				page = page + "</Container><Container><Input label='Your Twilio Phone Number' name='twilioPhoneNumber' /><Button action='selectNumber'>Use This Number</Button></Container><BR /><BR /><P>If you hate all the numbers above, you can make a new one that will be used</P><Button action='createInstead'>Create A Different Phone Number</Button></Page>"

				respond(res, page)
				client.Disconnect(context.TODO())
				return

			} else {
				// if there are no numbers - show them the form to create a phonenumber
				respond(res, "<Page><Container><H1>Create A Twilio Number</H1><Input label='Area Code of Your Desired Phone Number' name='areaCode' /></Container><Container><P>Note: you may be subject to charges from Twilio</P><Button action='createNumber'>Create Number</Button></Container></Page>")
				client.Disconnect(context.TODO())
				return
			}
		}

		respond(res, "<Page><Container><Link href='https://zeit.co'>Visit Twilio Console for Account Sid</Link><BR /><Link href='https://www.twilio.com/console/runtime/api-keys/create'>Visit Twilio API Console to Create an API Key</Link><Input label='Twilio API Key SID' name='APIKeySID' /><Input label='Twilio API Key Secret' name='apiKeySecret' /><Input label='Twilio Account SID' name='TwilioAccountSID' /></Container><Container><Button action='createUser'>Submit</Button></Container></Page>")
		client.Disconnect(context.TODO())
		return

	case "createInstead":
		respond(res, "<Page><Container><H1>Create A Twilio Number</H1><Input label='Area Code of Your Desired Phone Number' name='areaCode' /></Container><Container><P>Note: you may be subject to charges from Twilio</P><Button action='createNumber'>Create Number</Button></Container></Page>")
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
			page := "<Page><Container><H1>Select A Phone Number</H1><BR /><P><B>please type in one of the phone numbers listed below</B></P>"
			for _, number := range userNumbers {
				page = page + "<P>" + number + "</P><BR />"
			}
			page = page + "</Container><Container><Input label='Your Twilio Phone Number' name='twilioPhoneNumber' /><Button action='selectNumber'>Use This Number</Button></Container><BR /><BR /><P>If you hate all the numbers above, you can make a new one that will be used</P><Button action='createInstead'>Create A Different Phone Number</Button></Page>"

			respond(res, page)
			client.Disconnect(context.TODO())
			return
		} else {
			// if there are no numbers - show them the form to create a phonenumber
			respond(res, "<Page><Container><H1>Create A Twilio Number</H1><Input label='Area Code of Phone Number (3 digits)' name='areaCode' /></Container><Container><P>Note: you may be subject to charges from Twilio</P><Button action='createNumber'>Create Number</Button></Container></Page>")
			client.Disconnect(context.TODO())
			return
		}

	case "selectNumber":
		user := mymongo.FindUser(client, zeitPayload.User.Id)
		respond(res, NodeJS(user.TwilioAccountSID, zeitPayload.ClientState.TwilioPhoneNumbers, user.TwilioKeySID, user.TwilioKeySecret))
		client.Disconnect(context.TODO())
		return

	case "createNumber":

		user := mymongo.FindUser(client, zeitPayload.User.Id)

		newNumber := CreateNumber(user.TwilioKeySID, user.TwilioKeySecret, user.TwilioAccountSID, zeitPayload.ClientState.TwilioAreaCode)

		respond(res, NodeJS(user.TwilioAccountSID, newNumber, user.TwilioKeySID, user.TwilioKeySecret))
		client.Disconnect(context.TODO())
		return

	default:
		// just send out a default view
		respond(res, "<Page><P>There is a default view</P></Page>")
		client.Disconnect(context.TODO())
		return
	}
}

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
