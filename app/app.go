package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"github.com/adafatya/noobeeid-product_catalog/route"

	authrepository "github.com/adafatya/noobeeid-product_catalog/domain/auth"
	authhandler "github.com/adafatya/noobeeid-product_catalog/handler/auth"
	authrouter "github.com/adafatya/noobeeid-product_catalog/route/auth"
	authusecase "github.com/adafatya/noobeeid-product_catalog/use-case/auth"
)

type app struct {
	Db       *sqlx.DB
	RDb      *redis.Client
	FiberApp *fiber.App
}

func NewApp(db *sqlx.DB, rdb *redis.Client, fiberApp *fiber.App) app {
	return app{
		Db:       db,
		RDb:      rdb,
		FiberApp: fiberApp,
	}
}

func (app *app) Bootstrap() {
	// inject repository
	authRepository := authrepository.NewAuthRepository(app.Db, app.RDb)

	// inject use case
	authUseCae := authusecase.NewAuthUseCase(authRepository)

	// inject handler
	authHandler := authhandler.NewAuthHandler(authUseCae)

	// inject router
	authRouter := authrouter.NewAuthRouter(authHandler)

	// route
	route.Route(app.FiberApp, authRouter)
}
