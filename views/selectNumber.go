package views

func SelectNumber(userNumbers []string) string {
	languages := make([]string, 3)
	languages[0] = "Node.JS"
	languages[1] = "Python"
	languages[2] = "Php"
	page := `
	<Page>
		<Container>
			<H1>Select A Phone Number & A Programming Language</H1>
			<BR />
		</Container>
		<Container>
			<Select name="twilioPhoneNumber" value='` + userNumbers[0] + `' >`

	for _, number := range userNumbers {
		page = page + "<Option value='" + number + "' caption='" + number + "' />"
	}

	page = page + `
			</Select>
			<Select name="language" value='` + languages[0] + `'>`
	for _, language := range languages {
		page = page + "<Option value='" + language + "' caption='" + language + "' />"
	}
	page = page + `</Select>
			<Button action='selectNumber'>Use This Number & Language</Button>
		</Container>
		<BR />
		<P>If you hate all the numbers above, you can make a new one that will be used</P>
		<Button action='createInstead'>Create A Different Phone Number</Button>
		</Page>
	`

	return page
}
