package httpAccount

import "time"

type (
	AccountReqBody struct {
		Name   string `json:"name" validate:"required"`
		Cpf    string `json:"cpf" validate:"required"`
		Secret string `json:"secret" validate:"required"`
	}

	AcccountRespBody struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CPF       string    `json:"cpf"`
		Balance   float64   `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
	}
)
