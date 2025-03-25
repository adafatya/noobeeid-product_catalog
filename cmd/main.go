package main

import (
	"log"

	"github.com/adafatya/noobeeid-product_catalog/app"
	"github.com/adafatya/noobeeid-product_catalog/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	db := utils.ConnectDB()

	// migrate db
	migrate(db)

	rdb := utils.ConnectRedis()

	// create new fiber app
	fiberApp := fiber.New()

	// create app
	app := app.NewApp(db, rdb, fiberApp)

	// bootstrap app
	app.Bootstrap()

	// listen fiber app to port 8080
	app.FiberApp.Listen(":8080")
}
