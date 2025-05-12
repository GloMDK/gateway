package main

import (
	"gateway/clients/bank_clients/fast_bank"
	"gateway/clients/bank_clients/slow_bank"
	"gateway/clients/rates_client"
	"gateway/clients/transactions_client"
	"gateway/server"
	"gateway/service"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

func main() {
	app := fiber.New()
	appService := initService()
	appServer := server.New(appService)

	// app.Use(swagger.New())
	addRoutes(app, appServer)
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

func initService() server.Service {
	httpClient := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   60 * time.Second,
	}
	slowBankClient := slow_bank.New(nil)
	fastBankClient := fast_bank.New(nil)
	ratesClient := rates_client.New("http://rates:3001", httpClient, slowBankClient, fastBankClient)
	transactionsClient := transactions_client.New("http://transactions:3002", httpClient)

	return service.New(ratesClient, transactionsClient)
}
