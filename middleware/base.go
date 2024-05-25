package middleware

import (
	"github.com/LukaTchumburidze/simple-currency-converter/log"
	"github.com/gofiber/fiber/v2"
)

func AddMiddlewares(app *fiber.App) {
	app.Use(requestMiddleware)
	app.Use(responseMiddleware)

	log.Logger.Info("added middlewares")
}
