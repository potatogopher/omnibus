package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("user", func() {

	DefaultMedia(User)
	BasePath("/users")

	Action("show", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Retrieve user with given id. IDs 1 and 2 pre-exist in the system.")
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new user")
		Payload(func() {
			Member("givenName")
			Member("surname")
			Member("email")
			Member("password")
			Member("disabled")
			Required("email")
			Required("password")
		})
		Response(Created, "/user/[0-9]+")
	})

	Action("update", func() {
		Routing(
			PUT("/:userID"),
		)
		Description("Change user attributes")
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Payload(func() {
			Member("givenName")
			Member("surname")
			Member("email")
			Member("oldPassword")
			Member("newPassword")
			Member("disabled")
		})
		Response(NoContent)
		Response(NotFound)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
})
