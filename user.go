package main

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/goadesign/goa"

	"golang.org/x/crypto/scrypt"

	"goa-atlas/app"
	"goa-atlas/models"
)

var (
	// PasswordSaltBytes is the number of bytes a salt will be
	PasswordSaltBytes = 32

	// PasswordHashBytes is the number of bytes a hash will be
	PasswordHashBytes = 64
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

	u.Email = ctx.Payload.Email
	u.PasswordHash = hex.EncodeToString(hash)
	u.PasswordSalt = hex.EncodeToString(salt)

	user, err := udb.Add(ctx, &u)
	if err != nil {
		return err
	}

	ctx.ResponseData.Header().Set("Location", app.UserHref(user.ID))
	return nil
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	// TBD: implement
	return nil
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// TBD: implement
	res := &app.AtlasUser{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// TBD: implement
	return nil
}
