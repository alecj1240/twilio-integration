package handler

import (
	"net/url"
	"strings"
)

func NodeJS(accountSid string, phoneNumber string, apiKeySID string, apiKeySecret string) string {
	//str := encodeURIComponent("const accountSid = 'ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX';const authToken = 'your_auth_token';const client = require('twilio')(accountSid, authToken);client.messages.create%28%7Bfrom%3A%20%27%2B15017122661%27%2C%20body%3A%20%27body%27%2C%20to%3A%20%27%2B15558675310%27%7D%29.then(message => console.log(message.sid));")
	//return "<Page><Code>" + "<H1> Heillo </H1>" + "</Code></Page>"
	return "<Page><H1> Send an SMS message</H1><Code>" + "curl -X POST https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json --data-urlencode 'From=" + phoneNumber + "' --data-urlencode 'Body=Body' --data-urlencode 'To=" + phoneNumber + "' -u " + apiKeySID + ":" + apiKeySecret + "</Code></Page>"
}

func encodeURIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}
