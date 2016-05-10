package main

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/scrypt"

	"goa-blog/app"
	"goa-blog/models"
)

var (
	// PasswordSaltBytes is the number of bytes a salt will be
	PasswordSaltBytes = 32

	// PasswordHashBytes is the number of bytes a hash will be
	PasswordHashBytes = 64

	// ErrDatabaseError is the error returned when a db query fails.
	ErrDatabaseError = goa.NewErrorClass("db_error", 500)
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	u := models.User{}

	salt := make([]byte, PasswordSaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return err
	}

	hash, err := scrypt.Key([]byte(ctx.Payload.Password), salt, 1<<14, 8, 1, PasswordHashBytes)
	if err != nil {
		return err
	}

	u.GivenName = *ctx.Payload.GivenName
	u.Surname = *ctx.Payload.Surname
	u.Email = ctx.Payload.Email
	u.PasswordHash = hex.EncodeToString(hash)
	u.PasswordSalt = hex.EncodeToString(salt)

	user, err := udb.Add(ctx, &u)
	if err != nil {
		return ErrDatabaseError(err)
	}

	ctx.ResponseData.Header().Set("Location", app.UserHref(user.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	user, err := udb.Get(ctx.Context, ctx.UserID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}

	user.Disabled = true

	err = udb.Update(ctx, &user)
	if err != nil {
		return ErrDatabaseError(err)
	}

	return ctx.NoContent()
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	user, err := udb.OneUser(ctx.Context, ctx.UserID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}

	user.Href = app.UserHref(user.ID)
	return ctx.OK(user)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	user, err := udb.Get(ctx.Context, ctx.UserID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}

	user.GivenName = *ctx.Payload.GivenName
	user.Surname = *ctx.Payload.Surname
	user.Email = *ctx.Payload.Email

	err = udb.Update(ctx, &user)
	if err != nil {
		return ErrDatabaseError(err)
	}

	return ctx.NoContent()
}
