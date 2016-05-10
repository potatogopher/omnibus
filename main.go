package main

import (
	"fmt"
	"log"

	token "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"goa-blog/app"
	"goa-blog/models"
	"goa-blog/resources/config"
	"goa-blog/swagger"

	"gopkg.in/inconshreveable/log15.v2"
)

var db *gorm.DB
var logger log15.Logger
var udb *models.UserDB
var pdb *models.PostDB
var ErrDatabaseError = goa.NewErrorClass("db_error", 500)
var RSAPublicKey string

func main() {
	// Create service
	var err error
	url := fmt.Sprintf("dbname=gorma user=nicholasrucci sslmode=disable")
	fmt.Println(url)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	db.DropTable(&models.User{}, &models.Post{})
	db.AutoMigrate(&models.User{}, &models.Post{})

	udb = models.NewUserDB(*db)
	pdb = models.NewPostDB(*db)
	db.DB().SetMaxOpenConns(50)

	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Configure JWT Security
	publicKey, err := token.ParseRSAPublicKeyFromPEM([]byte(config.RSAPublicKey))
	if err != nil {
		log.Fatal(err)
	}
	app.ConfigureJWTSecurity(service, jwt.New(publicKey, nil))

	// Mount "user" controller
	c := NewUserController(service)
	app.MountUserController(service, c)
	// Mount "post" controller
	pc := NewPostController(service)
	app.MountPostController(service, pc)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
