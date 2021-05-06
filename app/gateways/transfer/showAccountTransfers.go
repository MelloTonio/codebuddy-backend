package httpTransfer

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h TransferHandler) ShowAccountTransfers(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "accountID")

	transfers, err := h.service.ShowAllAccountTransfers(accountID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusCreated, transfers)

}
