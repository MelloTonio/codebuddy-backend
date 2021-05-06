package httpAccount

import (
	"net/http"

	"github.com/mellotonio/desafiogo/app/domain/account"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var accountBody AccountReqBody

	err := response.Decode(r, accountBody)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.CreateAccount(&account.Account{
		Name:    accountBody.Name,
		Cpf:     accountBody.Cpf,
		Secret:  accountBody.Secret,
		Balance: 0,
	})

	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, accountBody)
}
