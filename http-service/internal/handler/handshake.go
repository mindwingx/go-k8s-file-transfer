package handler

import (
	"http-service/pkg/helper"
	"net/http"
)

type handshake struct {
	Message string `json:"msg"`
}

func Handshake(rw http.ResponseWriter, _ *http.Request) {
	data := handshake{
		Message: "connection established",
	}

	helper.JsonResponse(rw, http.StatusOK, data)
	return
}
