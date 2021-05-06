package httpAuth

import (
	"net/http"

	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var accessCredential access.Credential

	err := response.Decode(r, accessCredential)

	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	jwt, err := h.service.Authenticate(accessCredential)

	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	response.JSON(w, http.StatusOK, jwt)
}
