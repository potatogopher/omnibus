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
		Description("Retrieve user with given id.")
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

var _ = Resource("auth", func() {
	DefaultMedia(Authorize)
	BasePath("/auth")
	Action("token", func() {
		Routing(
			POST("/token"),
		)
		Description("Obtain an access token")
		Payload(func() {
			Member("email")
			Member("password")
			Required("email")
			Required("password")
		})
		Response(Created, func() {
			Media(Authorize)
		})
	})
})

var _ = Resource("post", func() {

	DefaultMedia(Post)
	BasePath("/posts")
	Security("jwt")

	Action("show", func() {
		Routing(
			GET("/:postID"),
		)
		Description("Retrieve post with given id.")
		Params(func() {
			Param("postID", Integer, "Post ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new post")
		Payload(func() {
			Member("title")
			Member("content")
			Required("title")
			Required("content")
		})
		Response(Created, "/posts/[0-9]+")
	})

	Action("update", func() {
		Routing(
			PUT("/:postID"),
		)
		Description("Change post attributes")
		Params(func() {
			Param("postID", Integer, "Post ID")
		})
		Payload(func() {
			Member("title")
			Member("content")
		})
		Response(NoContent)
		Response(NotFound)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:postID"),
		)
		Params(func() {
			Param("postID", Integer, "Post ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
})
