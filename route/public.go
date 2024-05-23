package route

import (
	"github.com/LukaTchumburidze/simple-currency-converter/controller"
	"github.com/LukaTchumburidze/simple-currency-converter/log"
	"github.com/gofiber/fiber/v2"
)

func AddPublicRoutes(fiber *fiber.App) {
	fiber.Post("/convert", controller.ConvertCurrency)

	log.Logger.Info("Added public routes")
}
