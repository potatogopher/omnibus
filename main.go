package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"goa-atlas/app"
	"goa-atlas/models"
	"goa-atlas/swagger"
	"gopkg.in/inconshreveable/log15.v2"
)

var db *gorm.DB
var logger log15.Logger
var udb *models.UserDB

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
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
