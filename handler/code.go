package handler

import "fmt"

func NodeJS(accountSid string, phoneNumber string, apiKeySID string, apiKeySecret string) string {
	str := `
	<Page>
		<Container>
			<H1> Node.JS Code </H1>
			<BR />
			<H2> This is the neccessary code to start sending texts with your Twilio phone number </H2>
			<P> Note, you will want to change the message parameters: to, from, and the body </P>
		</Container>
		<Container>
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
		</Container>
	</Page>
	`
	str = fmt.Sprintf(str, apiKeySID, apiKeySecret, phoneNumber, phoneNumber)
	return str
}

func Python(accountSid string, phoneNumber string, apiKeySID string, apiKeySecret string) string {
	str := `
	<Page>
		<Container>
			<H1> Python Code </H1>
			<BR />
			<H2> This is the neccessary code to start sending texts with your Twilio phone number </H2>
			<P> Note, you will want to change the message parameters: to, from, and the body </P>
		</Container>
		<Container>
			<Code>
				# Download the helper library from https://www.twilio.com/docs/python/install
				from twilio.rest import Client


				# Your Account Sid and Auth Token from twilio.com/console
				# DANGER! This is insecure. See http://twil.io/secure
				account_sid = '%s'
				auth_token = '%s'
				client = Client(account_sid, auth_token)

				message = client.messages \
					.create(
						body="Join Earth's mightiest heroes. Like Kevin Bacon.",
						from_='%s',
						to='%s'
				)

				print(message.sid)
			</Code>
		</Container>
	</Page>
	`
	str = fmt.Sprintf(str, apiKeySID, apiKeySecret, phoneNumber, phoneNumber)
	return str
}

func Php(accountSid string, phoneNumber string, apiKeySID string, apiKeySecret string) string {
	str := `
	<Page>
		<Container>
			<H1> Php Code </H1>
			<BR />
			<H2> This is the neccessary code to start sending texts with your Twilio phone number </H2>
			<P> Note, you will want to change the message parameters: to, from, and the body </P>
		</Container>
		<Container>
			<Code>
				&#60;?php
					require __DIR__ . '/vendor/autoload.php';
					use Twilio\Rest\Client;

					// Your Account SID and Auth Token from twilio.com/console
					$account_sid = '%s';
					$auth_token = '%s';
					// In production, these should be environment variables. E.g.:
					// $auth_token = $_ENV["TWILIO_ACCOUNT_SID"]

					// A Twilio number you own with SMS capabilities
					$twilio_number = "%s";

					$client = new Client($account_sid, $auth_token);
					$client->messages->create(
						// Where to send a text message (your cell phone?)
						'%s',
						array(
							'from' => $twilio_number,
							'body' => 'I sent this message in under 10 minutes!'
						)
					);
			</Code>
		</Container>
	</Page>
	`
	str = fmt.Sprintf(str, apiKeySID, apiKeySecret, phoneNumber, phoneNumber)
	return str
}
