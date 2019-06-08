package views

func SelectNumber(userNumbers []string) string {
	page := `
	<Page>
		<Container>
			<H1>Select A Phone Number</H1>
			<BR />
		</Container>
		<Container>
			<Select name="twilioPhoneNumber" value='` + userNumbers[0] + `' >`

	for _, number := range userNumbers {
		page = page + "<Option value='" + number + "' caption='" + number + "' />"
	}

	page = page + `
			</Select>
			<Button action='selectNumber'>Use This Number</Button>
		</Container>
		<BR />
		<P>If you hate all the numbers above, you can make a new one that will be used</P>
		<Button action='createInstead'>Create A Different Phone Number</Button>
		</Page>
	`

	return page
}
