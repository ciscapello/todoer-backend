package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReturnBadRequestError(writer http.ResponseWriter, statusCode int, key string, value string) {
	writer.WriteHeader(statusCode)
	writer.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp[key] = value
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, _ = writer.Write(jsonResp)
	return
}
