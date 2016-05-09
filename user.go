package main

import (
	"github.com/goadesign/goa"

	"goa-atlas/app"
	"goa-atlas/models"
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
	user, err := udb.Add(ctx, &models.User{})
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
