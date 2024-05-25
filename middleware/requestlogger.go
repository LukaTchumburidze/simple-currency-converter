package middleware

import (
	"github.com/LukaTchumburidze/simple-currency-converter/entity"
	"github.com/LukaTchumburidze/simple-currency-converter/repository"
	"github.com/LukaTchumburidze/simple-currency-converter/request"
	"github.com/gofiber/fiber/v2"
)

const RequestIDKey = "request"
const CurrencyKey = "currency"

func requestMiddleware(ctx *fiber.Ctx) error {
	var currency request.Currency
	err := currency.BuildFromFiberCtx(ctx)
	if err != nil {
		return err
	}

	storedRequest, err := repository.Aggregator.RequestRepository.StoreRequest(entity.Request{
		Headers: string(ctx.Request().Header.Header()),
		Body:    string(ctx.Request().Body()),
	})
	if err != nil {
		return err
	}
	ctx.Locals(RequestIDKey, storedRequest.ID)
	ctx.Locals(CurrencyKey, currency)

	return ctx.Next()
}
