//************************************************************************//
// API "Atlas": Application Contexts
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

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(ctx context.Context, service *goa.Service) (*CreateUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := CreateUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// createUserPayload is the user create action payload.
type createUserPayload struct {
	// Flag for if the user is disabled or not
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// Email of user
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Given name of user
	GivenName *string `json:"givenName,omitempty" xml:"givenName,omitempty"`
	Password  *string `json:"password,omitempty" xml:"password,omitempty"`
	// Surname of user
	Surname *string `json:"surname,omitempty" xml:"surname,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createUserPayload) Validate() (err error) {
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}

	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return err
}

// Publicize creates CreateUserPayload from createUserPayload
func (payload *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if payload.Disabled != nil {
		pub.Disabled = payload.Disabled
	}
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.GivenName != nil {
		pub.GivenName = payload.GivenName
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	if payload.Surname != nil {
		pub.Surname = payload.Surname
	}
	return &pub
}

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	// Flag for if the user is disabled or not
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// Email of user
	Email string `json:"email" xml:"email"`
	// Given name of user
	GivenName *string `json:"givenName,omitempty" xml:"givenName,omitempty"`
	Password  string  `json:"password" xml:"password"`
	// Surname of user
	Surname *string `json:"surname,omitempty" xml:"surname,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateUserPayload) Validate() (err error) {
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}

	if err2 := goa.ValidateFormat(goa.FormatEmail, payload.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, payload.Email, goa.FormatEmail, err2))
	}
	return err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteUserContext provides the user delete action context.
type DeleteUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	UserID  int
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(ctx context.Context, service *goa.Service) (*DeleteUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DeleteUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	UserID  int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(ctx context.Context, service *goa.Service) (*ShowUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(r *AtlasUser) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/atlas.user")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OKTiny(r *AtlasUserTiny) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/atlas.user")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	UserID  int
	Payload *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(ctx context.Context, service *goa.Service) (*UpdateUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := UpdateUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// updateUserPayload is the user update action payload.
type updateUserPayload struct {
	// Flag for if the user is disabled or not
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// Email of user
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Given name of user
	GivenName   *string `json:"givenName,omitempty" xml:"givenName,omitempty"`
	NewPassword *string `json:"newPassword,omitempty" xml:"newPassword,omitempty"`
	OldPassword *string `json:"oldPassword,omitempty" xml:"oldPassword,omitempty"`
	// Surname of user
	Surname *string `json:"surname,omitempty" xml:"surname,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateUserPayload) Validate() (err error) {
	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return err
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (payload *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if payload.Disabled != nil {
		pub.Disabled = payload.Disabled
	}
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.GivenName != nil {
		pub.GivenName = payload.GivenName
	}
	if payload.NewPassword != nil {
		pub.NewPassword = payload.NewPassword
	}
	if payload.OldPassword != nil {
		pub.OldPassword = payload.OldPassword
	}
	if payload.Surname != nil {
		pub.Surname = payload.Surname
	}
	return &pub
}

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	// Flag for if the user is disabled or not
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// Email of user
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Given name of user
	GivenName   *string `json:"givenName,omitempty" xml:"givenName,omitempty"`
	NewPassword *string `json:"newPassword,omitempty" xml:"newPassword,omitempty"`
	OldPassword *string `json:"oldPassword,omitempty" xml:"oldPassword,omitempty"`
	// Surname of user
	Surname *string `json:"surname,omitempty" xml:"surname,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateUserPayload) Validate() (err error) {
	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
