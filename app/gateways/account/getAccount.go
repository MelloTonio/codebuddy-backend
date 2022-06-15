package httpAccount

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "accountID")

	ac, err := h.service.GetAccount(accountID)

	if err != nil {
		response.Error(w, http.StatusNotFound, err)
		return
	}

	response.JSON(w, http.StatusCreated, ac)
}
