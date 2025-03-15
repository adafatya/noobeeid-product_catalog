package main

import (
	"github.com/adafatya/noobeeid-product_catalog/utils"
)

func main() {
	// connect to database
	db := utils.ConnectDB()

	// migrate db
	migrate(db)

}
