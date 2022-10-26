package main

import (
	"open-funding/config"
	"open-funding/exceptions"

	database "open-funding/utils/database/postgre"

	"open-funding/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type CustomValidation struct {
	validator *validator.Validate
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: exceptions.ExceptionHandler,
	})
	app.Use(recover.New())

	cfg := config.GetConfig()
	db := database.InitDB(cfg)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	routes.InitRoutes(app, db, cfg)

	err := app.Listen(":" + cfg.SERVER_PORT)

	if err != nil {
		panic(err)
	}
}
