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
		Attribute("given_name", String, "Given name of user")
		Attribute("surname", String, "Surname of user")
		Attribute("email", String, "Email of user", func() {
			Format("email")
		})
		Attribute("password_hash", String, "Hashed password of user")
		Attribute("password_salt", String, "Salted password of user")
		Attribute("disabled", Boolean, "Flag for if the user is disabled or not")

		Required("id", "href", "email", "password_salt", "password_hash", "disabled")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("email")
		Attribute("given_name")
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
