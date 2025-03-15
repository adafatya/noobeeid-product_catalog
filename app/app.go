package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type app struct {
	Db       *sqlx.DB
	FiberApp *fiber.App
}

func NewApp(db *sqlx.DB, fiberApp *fiber.App) app {
	return app{
		Db:       db,
		FiberApp: fiberApp,
	}
}

func (app app) Bootstrap() {

}
