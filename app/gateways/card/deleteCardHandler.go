package httpCard

import (
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h CardHandler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("cardID")

	err := h.cardUC.Delete(userID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, 200, `"success": true`)
}
