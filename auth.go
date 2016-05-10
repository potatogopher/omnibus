package main

import (
	"github.com/goadesign/goa"
	"goa-blog/app"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
}

// Callback runs the callback action.
func (c *AuthController) Callback(ctx *app.CallbackAuthContext) error {
	// TBD: implement
	return nil
}

// Oauth runs the oauth action.
func (c *AuthController) Oauth(ctx *app.OauthAuthContext) error {
	// TBD: implement
	res := &app.Authorize{}
	return ctx.OK(res)
}

// Refresh runs the refresh action.
func (c *AuthController) Refresh(ctx *app.RefreshAuthContext) error {
	// TBD: implement
	return nil
}

// Token runs the token action.
func (c *AuthController) Token(ctx *app.TokenAuthContext) error {
	// TBD: implement
	return nil
}
