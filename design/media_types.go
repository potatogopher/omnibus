package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// User is the user resource media type
var User = MediaType("application/vnd.user+json", func() {
	Description("A user account")
	Attributes(func() {
		Attribute("id", Integer, "ID of user")
		Attribute("href", String, "API href of the user")
		Attribute("given_name", String, "Given name of user")
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

// Authorize is the authorize resource media type
var Authorize = MediaType("application/vnd.authorize+json", func() {
	Description("Token authorization response")
	Attributes(func() {
		Attribute("access_token", String, "access token", func() {})
		Attribute("expires_in", String, "Time to expiration in seconds", func() {})
		Attribute("token_type", String, "type of token", func() {})
	})

	View("default", func() {
		Attribute("access_token")
		Attribute("expires_in")
		Attribute("token_type")
	})
})

// Post is the post resource media type
var Post = MediaType("application/vnd.post+json", func() {
	Description("A blog post")
	Attributes(func() {
		Attribute("id", Integer, "ID of post")
		Attribute("href", String, "API href of the post")
		Attribute("title", String, "Title of the blog post")
		Attribute("content", String, "Content of the blog post")

		Required("id", "href", "title", "content")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("title")
		Attribute("content")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})
