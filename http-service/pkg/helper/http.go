package helper

import (
	"encoding/json"
	"net/http"
	"time"
)

type response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
	Time  string      `json:"time"`
}

func JsonResponse(rw http.ResponseWriter, statusCode int, payload interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)

	resp := new(response)
	resp.Time = time.Now().Format(time.RFC3339)

	if statusCode >= 200 && statusCode < 400 {
		resp.Data = payload
	} else {
		resp.Error = payload
	}

	_ = json.NewEncoder(rw).Encode(resp)
}
