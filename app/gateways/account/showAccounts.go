package httpAccount

import (
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h AccountHandler) ShowAccounts(w http.ResponseWriter, r *http.Request) {

	ac, err := h.service.ShowAccounts()

	if err != nil {
		response.Error(w, http.StatusNotFound, err)
		return
	}

	response.JSON(w, http.StatusFound, ac)
}
