package status

import (
	"encoding/json"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, stat := Http400(err.Error(), nil)
	statBytes, _ := json.Marshal(stat)
	w.Write(statBytes)
}
