package main

import (
	"gateway/gateway/server"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"log"
)

func main() {
	app := fiber.New()
	s := server.New(nil)

	app.Use(swagger.New())
	addRoutes(app, s)
	initLogger(app)

	log.Fatal(app.Listen(":3000"))
}

func addRoutes(app *fiber.App, s *server.Server) {
	app.Post("/pay", s.Pay)
	app.Get("/pay_status", s.PayStatus)
}

func initLogger(app *fiber.App) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
}
