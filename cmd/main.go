package main

import (
	"github.com/adafatya/noobeeid-product_catalog/app"
	"github.com/adafatya/noobeeid-product_catalog/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// connect to database
	db := utils.ConnectDB()

	// migrate db
	migrate(db)

	// create new fiber app
	fiberApp := fiber.New()

	// create app
	app := app.NewApp(db, fiberApp)

	// bootstrap app
	app.Bootstrap()

	// listen fiber app to port 8080
	app.FiberApp.Listen(":8080")
}
