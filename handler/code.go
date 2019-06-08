package handler

import "fmt"

func NodeJS(accountSid string, phoneNumber string, apiKeySID string, apiKeySecret string) string {
	str := `
	<Page>
		<Code>
			// Download the helper library from https://www.twilio.com/docs/node/install
			// Your Account Sid and Auth Token from twilio.com/console
			// DANGER! This is insecure. See http://twil.io/secure
			const accountSid = '%s';
			const authToken = '%s';
			const client = require('twilio')(accountSid, authToken);

			client.messages
				.create(&#123;
					body: 'This is the ship that made the Kessel Run in fourteen parsecs?',
					from: '%s',
					to: '%s'
			&#125;)
			.then(message => console.log(message.sid));
		</Code>
	</Page>
	`
	str = fmt.Sprintf(str, apiKeySID, apiKeySecret, phoneNumber, phoneNumber)
	return str
}
