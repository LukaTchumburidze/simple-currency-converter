package middleware

import (
	"fmt"
	"github.com/LukaTchumburidze/simple-currency-converter/entity"
	"github.com/LukaTchumburidze/simple-currency-converter/repository"
	"github.com/LukaTchumburidze/simple-currency-converter/request"
	"github.com/gofiber/fiber/v2"
)

const requestKey = "request"

func requestMiddleware(ctx *fiber.Ctx) error {
	var currency request.Currency
	err := currency.BuildFromFiberCtx(ctx)
	if err != nil {
		return err
	}

	headers := fmt.Sprintf("%s", ctx.Request().Header.Header())
	body := fmt.Sprintf("%s", ctx.Request().Body())

	storedRequest, err := repository.Aggregator.RequestRepository.StoreRequest(entity.Request{
		Headers: headers,
		Body:    body,
	})
	if err != nil {
		return err
	}
	ctx.Locals(requestKey, storedRequest)

	return ctx.Next()
}
