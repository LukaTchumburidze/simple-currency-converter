package response

type Currency struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
}
