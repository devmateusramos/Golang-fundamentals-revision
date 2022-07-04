package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatalln(err)
		}
	}
}

func Err(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Err string `json:"err"`
	}{
		Err: err.Error(),
	})
}
