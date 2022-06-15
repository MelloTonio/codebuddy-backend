package httpCard

import (
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h CardHandler) GetAllCards(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")

	cards, err := h.cardUC.Get(userID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, cards)
}
