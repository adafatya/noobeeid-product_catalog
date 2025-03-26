package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"github.com/adafatya/noobeeid-product_catalog/route"

	authrepository "github.com/adafatya/noobeeid-product_catalog/domain/auth"
	categoryrepository "github.com/adafatya/noobeeid-product_catalog/domain/category"
	authhandler "github.com/adafatya/noobeeid-product_catalog/handler/auth"
	categoryhandler "github.com/adafatya/noobeeid-product_catalog/handler/category"
	authrouter "github.com/adafatya/noobeeid-product_catalog/route/auth"
	categoryrouter "github.com/adafatya/noobeeid-product_catalog/route/category"
	authusecase "github.com/adafatya/noobeeid-product_catalog/use-case/auth"
	categoryusecase "github.com/adafatya/noobeeid-product_catalog/use-case/category"
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
	categoryRepository := categoryrepository.NewCategoryRepository(app.Db)

	// inject use case
	authUseCae := authusecase.NewAuthUseCase(authRepository)
	categoryUseCase := categoryusecase.NewCategoryUseCase(categoryRepository)

	// inject handler
	authHandler := authhandler.NewAuthHandler(authUseCae)
	categoryHandler := categoryhandler.NewCategoryHandler(categoryUseCase)

	// inject router
	authRouter := authrouter.NewAuthRouter(authHandler)
	categoryRouter := categoryrouter.NewCategoryRouter(categoryHandler)

	// route
	route.Route(app.FiberApp, authRouter, categoryRouter)
}
