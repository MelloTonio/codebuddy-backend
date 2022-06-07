package httpCard

import (
	"net/http"

	"github.com/mellotonio/desafiogo/app/domain/card"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

type CardReqBody struct {
	Owner    string `json:"owner"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (h CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	var reqBody card.Card

	err := response.Decode(r, &reqBody)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err = h.cardUC.Store(&reqBody)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, reqBody)
}
