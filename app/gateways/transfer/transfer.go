package httpTransfer

import (
	"net/http"
	"strconv"

	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/mellotonio/desafiogo/app/infra/utils"
)

func (h TransferHandler) Transfer(w http.ResponseWriter, r *http.Request) {

	// Must be logged in to perform a transaction
	accDescription, err := utils.ExtractTokenMetadata(r)

	if err != nil {
		err := errors.ErrUnauthorized
		response.Error(w, http.StatusUnauthorized, err)
	}

	accountDestinationID := r.URL.Query().Get("accountID_destination")
	amount := r.URL.Query().Get("amount")

	amountInt, err := strconv.Atoi(amount)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	transfer, err := h.service.Transfer(accDescription.AccountID, accountDestinationID, amountInt)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusCreated, transfer)

}
