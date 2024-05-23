package request

import (
	"github.com/gofiber/fiber/v2"
)

type Currency struct {
	BaseCurrency   string  `binding:"base_currency" validate:"iso4217"`
	TargetCurrency string  `binding:"target_currency" validate:"iso4217"`
	Amount         float64 `json:"amount" validate:"required,gt=0"`
}

func (c Currency) BuildFromFiberCtx(ctx *fiber.Ctx) error {
	err := ctx.ParamsParser(c)
	if err != nil {
		return err
	}
	err = ctx.BodyParser(c)
	if err != nil {
		return err
	}
	return validator.Struct(c)
}
