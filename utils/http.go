package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(
	w http.ResponseWriter,
	httpCode int,
	message string,
	payload interface{},
) {
	respPayload := map[string]interface{}{
		"stat_msg": message,
		"data":     payload,
	}

	response, _ := json.Marshal(respPayload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	_, _ = w.Write(response)
}
