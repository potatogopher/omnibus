package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var User = MediaType("application/atlas.user+json", func() {
	Description("A user account")
	Attributes(func() {
		Attribute("id", Integer, "ID of user")
		Attribute("href", String, "API href of the user")
		Attribute("givenName", String, "Given name of user")
		Attribute("surname", String, "Surname of user")
		Attribute("email", String, "Email of user", func() {
			Format("email")
		})
		Attribute("passwordHash", String, "Hashed password of user")
		Attribute("passwordSalt", String, "Salted password of user")
		Attribute("disabled", Boolean, "Flag for if the user is disabled or not")

		Required("id", "href", "email", "passwordSalt", "passwordHash", "disabled")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("email")
		Attribute("givenName")
		Attribute("surname")
		Attribute("disabled")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("href")
		Attribute("email")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})
