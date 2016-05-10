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

// Post media type.
//
// Identifier: application/vnd.post+json
type Post struct {
	// Content of the blog post
	Content string `json:"content" xml:"content"`
	// API href of the post
	Href string `json:"href" xml:"href"`
	// ID of post
	ID int `json:"id" xml:"id"`
	// Title of the blog post
	Title string `json:"title" xml:"title"`
}

// Validate validates the Post media type instance.
func (mt *Post) Validate() (err error) {
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

// PostLink media type.
//
// Identifier: application/vnd.post+json
type PostLink struct {
	// API href of the post
	Href string `json:"href" xml:"href"`
	// ID of post
	ID int `json:"id" xml:"id"`
}

// Validate validates the PostLink media type instance.
func (mt *PostLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return err
}

// User media type.
//
// Identifier: application/vnd.user+json
type User struct {
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

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
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

// UserLink media type.
//
// Identifier: application/vnd.user+json
type UserLink struct {
	// API href of the user
	Href string `json:"href" xml:"href"`
	// ID of user
	ID int `json:"id" xml:"id"`
}

// Validate validates the UserLink media type instance.
func (mt *UserLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return err
}

// UserTiny media type.
//
// Identifier: application/vnd.user+json
type UserTiny struct {
	// Email of user
	Email string `json:"email" xml:"email"`
	// API href of the user
	Href string `json:"href" xml:"href"`
	// ID of user
	ID int `json:"id" xml:"id"`
}

// Validate validates the UserTiny media type instance.
func (mt *UserTiny) Validate() (err error) {
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
