package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"http-service/api"
	"http-service/pkg/helper"
	"io"
	"net/http"
)

const uploadSize = 1 * 1024 * 1024 // up to 1MB

func Upload(client *grpc.ClientConn) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(rw, r.Body, uploadSize)
		if err := r.ParseMultipartForm(uploadSize); err != nil {
			helper.JsonResponse(rw, http.StatusUnprocessableEntity, "file too large. maximum size is 1MB")
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			helper.JsonResponse(rw, http.StatusBadRequest, "invalid file")
			return
		}

		defer file.Close()

		if !legalFileFormat(header.Header.Get("Content-Type")) {
			helper.JsonResponse(rw, http.StatusUnprocessableEntity, "invalid file format(only jpeg,png,jpg)")
			return
		}

		imageData, readErr := io.ReadAll(file)
		if readErr != nil {
			err = fmt.Errorf("[http] failed to read image data: %w", readErr)
			return
		}

		gc := api.NewUploadClient(client)

		uploadReq := api.UploadReq{
			File: &api.File{
				ContentType:  header.Header.Get("Content-Type"),
				ContentBytes: imageData,
			},
		}

		uploadResp, err := gc.UploadImage(context.Background(), &uploadReq)
		if err != nil {
			helper.JsonResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		helper.JsonResponse(rw, http.StatusCreated, uploadResp.Message)
		return
	}
}

// HELPERS

func legalFileFormat(fileMimeType string) bool {
	legalFormats := map[string]bool{"image/png": true, "image/jpg": true, "image/jpeg": true}
	return legalFormats[fileMimeType]
}
