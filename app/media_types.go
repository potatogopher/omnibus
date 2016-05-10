//************************************************************************//
// API "rucci.io": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/goa-atlas
// --design=goa-atlas/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// RucciPost media type.
//
// Identifier: application/rucci.post+json
type RucciPost struct {
	// Content of the blog post
	Content string `json:"content" xml:"content"`
	// API href of the post
	Href string `json:"href" xml:"href"`
	// ID of post
	ID int `json:"id" xml:"id"`
	// Title of the blog post
	Title string `json:"title" xml:"title"`
}

// Validate validates the RucciPost media type instance.
func (mt *RucciPost) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Content == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "content"))
	}

	return err
}

// RucciPostLink media type.
//
// Identifier: application/rucci.post+json
type RucciPostLink struct {
	// API href of the post
	Href string `json:"href" xml:"href"`
	// ID of post
	ID int `json:"id" xml:"id"`
}

// Validate validates the RucciPostLink media type instance.
func (mt *RucciPostLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return err
}

// RucciUser media type.
//
// Identifier: application/rucci.user+json
type RucciUser struct {
	// Flag for if the user is disabled or not
	Disabled bool `json:"disabled" xml:"disabled"`
	// Email of user
	Email string `json:"email" xml:"email"`
	// Given name of user
	GivenName *string `json:"givenName,omitempty" xml:"givenName,omitempty"`
	// API href of the user
	Href string `json:"href" xml:"href"`
	// ID of user
	ID int `json:"id" xml:"id"`
	// Surname of user
	Surname *string `json:"surname,omitempty" xml:"surname,omitempty"`
}

// Validate validates the RucciUser media type instance.
func (mt *RucciUser) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}

	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	return err
}

// RucciUserLink media type.
//
// Identifier: application/rucci.user+json
type RucciUserLink struct {
	// API href of the user
	Href string `json:"href" xml:"href"`
	// ID of user
	ID int `json:"id" xml:"id"`
}

// Validate validates the RucciUserLink media type instance.
func (mt *RucciUserLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return err
}

// RucciUserTiny media type.
//
// Identifier: application/rucci.user+json
type RucciUserTiny struct {
	// Email of user
	Email string `json:"email" xml:"email"`
	// API href of the user
	Href string `json:"href" xml:"href"`
	// ID of user
	ID int `json:"id" xml:"id"`
}

// Validate validates the RucciUserTiny media type instance.
func (mt *RucciUserTiny) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}

	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	return err
}
