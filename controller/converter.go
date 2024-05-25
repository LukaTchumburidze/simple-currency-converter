package controller

import (
	"github.com/LukaTchumburidze/simple-currency-converter/log"
	"github.com/LukaTchumburidze/simple-currency-converter/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ConvertCurrency(ctx *fiber.Ctx) error {
	log.Logger.WithFields(logrus.Fields{
		middleware.RequestIDKey: ctx.Locals(middleware.RequestIDKey),
		middleware.CurrencyKey:  ctx.Locals(middleware.CurrencyKey),
	}).Info("ConvertCurrency")
	ctx.Response().SetStatusCode(http.StatusOK)
	return nil
}
