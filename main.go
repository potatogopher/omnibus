package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"goa-blog/app"
	"goa-blog/models"
	"goa-blog/swagger"
	"gopkg.in/inconshreveable/log15.v2"
)

var db *gorm.DB
var logger log15.Logger
var udb *models.UserDB
var pdb *models.PostDB
var ErrDatabaseError = goa.NewErrorClass("db_error", 500)

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
