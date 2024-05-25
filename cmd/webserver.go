package cmd

import (
	"github.com/LukaTchumburidze/simple-currency-converter/log"
	"github.com/LukaTchumburidze/simple-currency-converter/middleware"
	"github.com/LukaTchumburidze/simple-currency-converter/route"
	"github.com/gofiber/fiber/v2"
)

func Execute() {
	app := fiber.New()

	middleware.AddMiddlewares(app)
	route.AddPublicRoutes(app)

	log.Logger.Info("Starting server on port 3000")
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
