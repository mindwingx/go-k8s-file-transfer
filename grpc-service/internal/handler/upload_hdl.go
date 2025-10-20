package handler

import (
	"context"
	"fmt"
	"grpc-service/api"
	"grpc-service/pkg/helper"
	"grpc-service/pkg/img"
	"os"
	"time"
)

type UploadHandler struct {
	api.UnimplementedUploadServer
}

func (s *UploadHandler) UploadImage(_ context.Context, req *api.UploadReq) (res *api.UploadResp, err error) {
	res = &api.UploadResp{}

	ext := img.MimeToExtension(req.File.GetContentType())
	filename := fmt.Sprintf("%s/assets/img_%d.%s", helper.Root(), time.Now().UnixNano(), ext)

	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("file create err: %s\n", err)
		return
	}

	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Printf("conn. close err: %s\n", err)
		}
	}()

	_, err = f.Write(req.File.ContentBytes)
	if err != nil {
		fmt.Printf("file write err: %s\n", err)
		return
	}

	fmt.Printf("image uploaded: %s\n", filename)

	res.Message = "api successful"
	return
}
