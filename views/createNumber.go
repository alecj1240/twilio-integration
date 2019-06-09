package views

func CreateNumber() string {
	page := `
		<Page>
			<Container>
				<H1>Create A Twilio Number</H1>
				<Input label='Area Code of Phone Number (3 digits)' name='areaCode' />
			</Container>
			<Container>
				<P>Note: you may be subject to charges from Twilio</P>
				<Button action='createNumber'>Create Number</Button>
			</Container>
		</Page>
	`
	return page
}