package middleware

import "github.com/gofiber/fiber/v2"

func AddMiddlewares(app *fiber.App) {
	app.Use(requestMiddleware)
	app.Use(responseMiddleware)
}
