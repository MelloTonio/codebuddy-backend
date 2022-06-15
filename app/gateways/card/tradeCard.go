package httpCard

import (
	"fmt"
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h CardHandler) TradeCard(w http.ResponseWriter, r *http.Request) {
	ownerID := r.URL.Query().Get("ownerID")
	cardID := r.URL.Query().Get("cardID")
	receiverID := r.URL.Query().Get("receiverID")

	fmt.Println(ownerID, cardID, receiverID)

	err := h.cardUC.TradeCard(cardID, ownerID, receiverID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, `"success": true"`)
}
