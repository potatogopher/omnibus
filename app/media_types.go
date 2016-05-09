//************************************************************************//
// API "Atlas": Application Media Types
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

// AtlasUser media type.
//
// Identifier: application/atlas.user+json
type AtlasUser struct {
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

// Validate validates the AtlasUser media type instance.
func (mt *AtlasUser) Validate() (err error) {
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

// AtlasUserLink media type.
//
// Identifier: application/atlas.user+json
type AtlasUserLink struct {
	// API href of the user
	Href string `json:"href" xml:"href"`
	// ID of user
	ID int `json:"id" xml:"id"`
}

// Validate validates the AtlasUserLink media type instance.
func (mt *AtlasUserLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return err
}

// AtlasUserTiny media type.
//
// Identifier: application/atlas.user+json
type AtlasUserTiny struct {
	// Email of user
	Email string `json:"email" xml:"email"`
	// API href of the user
	Href string `json:"href" xml:"href"`
	// ID of user
	ID int `json:"id" xml:"id"`
}

// Validate validates the AtlasUserTiny media type instance.
func (mt *AtlasUserTiny) Validate() (err error) {
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
