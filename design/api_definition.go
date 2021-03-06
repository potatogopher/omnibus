package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the cellar application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("Omnibus", func() {
	Title("Omnibus")
	Description("Blog for Nick Rucci")
	Contact(func() {
		Name("Nick Rucci")
		Email("nick@rucci.io")
		URL("http://rucci.io")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/nicholasrucci/omnibus/blob/master/LICENSE")
	})
	/*
	   Docs(func() {
	   		Description("goa guide")
	   		URL("http://goa.design/getting-started.html")
	   	})
	*/
	Host("localhost")
	Scheme("http")
	BasePath("/")

	Origin("http://swagger.localhost", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		MaxAge(600)
		Credentials()
	})
	JWTSecurity("jwt", func() {
		Header("Authorization")
		TokenURL("<a href='http://example.com/token'>http://example.com/token</a>")
		Scope("post:write", "Write to the system")
		Scope("post:read", "Read anything in there")
	})
	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
