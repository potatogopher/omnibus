package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("Atlas", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		Model("User", func() {
			RendersTo(User)
			Description("Atlas User")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("givenName", gorma.String)
			Field("surname", gorma.String)
			Field("email", gorma.String)
			Field("passwordHash", gorma.String)
			Field("passwordSalt", gorma.String)
			Field("disabled", gorma.Boolean)
			HasMany("Posts", "Post")
		})

		Model("Post", func() {
			BuildsFrom(func() {
				Payload("post", "create")
				Payload("post", "update")
			})
			RendersTo(Post)
			Description("Blog Post")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("title", gorma.String)
			Field("content", gorma.String)
			BelongsTo("User")
		})
	})
})
