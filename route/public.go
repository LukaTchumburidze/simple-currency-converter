package route

import (
	"github.com/LukaTchumburidze/simple-currency-converter/log"
	"github.com/gofiber/fiber/v2"
)

func AddPublicRoutes(fiber *fiber.App) {
	// TODO
	fiber.Post("/convert", nil)

	log.Logger.Info("Added public routes")
}
