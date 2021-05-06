package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Json Response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Decode(req *http.Request, body interface{}) error {
	return json.NewDecoder(req.Body).Decode(body)
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Erro    string `json:"error"`
		Success bool
	}{
		Erro:    err.Error(),
		Success: false,
	})

}
