package httpTransfer

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h TransferHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	accountOriginID := chi.URLParam(r, "accountID_origin")
	accountDestinationID := chi.URLParam(r, "accountID_destination")
	amount := chi.URLParam(r, "amount")

	amountInt, err := strconv.Atoi(amount)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	transfer, err := h.service.Transfer(accountOriginID, accountDestinationID, amountInt)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusCreated, transfer)

}
