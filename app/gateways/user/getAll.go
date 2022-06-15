package httpUser

import (
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUC.GetAll()
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, users)
}
