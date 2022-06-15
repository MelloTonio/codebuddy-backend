package httpUser

import (
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	nickName := r.URL.Query().Get("nickname")

	user, err := h.userUC.Get(nickName)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}
