package middleware

import (
	"encoding/json"
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

	headers, err := json.Marshal(ctx.Response().Header.Header())
	if err != nil {
		return err
	}
	body, err := json.Marshal(ctx.Response().Body())
	if err != nil {
		return err
	}
	requestId, err := strconv.ParseUint(fmt.Sprintf("%s", ctx.Locals(requestKey)), 10, 32)
	if err != nil {
		return err
	}

	_, err = repository.Aggregator.ResponseRepository.StoreResponse(entity.Response{
		RequestID: uint(requestId),
		Headers:   string(headers),
		Body:      string(body),
	})
	if err != nil {
		return err
	}

	return ctx.Next()
}
