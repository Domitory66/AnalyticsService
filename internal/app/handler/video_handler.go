package handler

import (
	"AnalyticsService/internal/app/api/image/prot"
	"AnalyticsService/internal/app/models"
	"context"
	"fmt"
)

type VideoHandler struct {
	prot.ImageWorkerClient
}

func (vh *VideoHandler) FindGesture(ctx context.Context, sendframe []byte, camera *models.Camera) (frame []byte, typeGesture string, err error) {
	Response := make(chan prot.MsgImageResponse)
	go func() {
		ctx := context.TODO()
		resp, err := vh.SearchGesture(ctx, &prot.MsgImageRequest{UserId: 0, CameraName: camera.Name, Image: sendframe})
		fmt.Println(err)
		Response <- *resp
	}()
	val := <-Response
	return val.Image, val.Label, nil
}
