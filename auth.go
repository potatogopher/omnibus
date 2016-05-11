package main

import (
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"

	"goa-blog/app"
	"goa-blog/resources/config"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
}

// Token runs the token action.
func (c *AuthController) Token(ctx *app.TokenAuthContext) error {
	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.RSAPrivateKey))
	if err != nil {
		return err
	}
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return err
	}

	authToken := app.Authorize{}
	authToken.AccessToken = &tokenStr

	res, err := json.Marshal(authToken)
	if err != nil {
		return err
	}

	return ctx.OK(res)
}
