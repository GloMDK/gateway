package main

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"log"
	"rates/cache_client"
	"rates/server"
)

func main() {
	app := fiber.New()
	cache := cache_client.New()
	appServer := server.New(cache)

	addRoutes(app, appServer)
	initLogger(app)

	log.Fatal(app.Listen(":3001"))
}

func addRoutes(app *fiber.App, s *server.Server) {
	app.Get("/param", s.GetParam)
	app.Patch("/param", s.UpdateParam)
	app.Post("/choose_bank_name", s.ChooseBankName)
}

func initLogger(app *fiber.App) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
}
