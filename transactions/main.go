package main

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"log"
	"transactions/server"
)

func main() {
	app := fiber.New()
	appServer, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	addRoutes(app, appServer)
	initLogger(app)

	log.Fatal(app.Listen(":3002"))
}

func addRoutes(app *fiber.App, s *server.Server) {
	app.Get("/one/+", s.Get)
	app.Post("/", s.Create)
	app.Patch("/+", s.Update)
	app.Get("/all", s.GetAll)
}
func initLogger(app *fiber.App) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
}
