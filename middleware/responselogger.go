package middleware

import (
	"fmt"
	"github.com/LukaTchumburidze/simple-currency-converter/entity"
	"github.com/LukaTchumburidze/simple-currency-converter/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func responseMiddleware(ctx *fiber.Ctx) error {
	err := ctx.Next()
	if err != nil {
		return err
	}

	requestId, err := strconv.ParseUint(fmt.Sprintf("%d", ctx.Locals(RequestIDKey)), 10, 32)
	if err != nil {
		return err
	}

	_, err = repository.Aggregator.ResponseRepository.StoreResponse(entity.Response{
		Status:    ctx.Response().StatusCode(),
		RequestID: uint(requestId),
		Headers:   string(ctx.Response().Header.Header()),
		Body:      string(ctx.Response().Body()),
	})
	if err != nil {
		return err
	}

	return nil
}
