package handler

import (
	"AnalyticsService/internal/app/models"
	"context"
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

type CameraHandler struct {
}

func (ch *CameraHandler) FindCamera(ctx context.Context, name string, port string, ip string, protocol string, filename string) (*models.Camera, error) {
	cap, err := gocv.VideoCaptureFile(protocol + "://" + ip + ":" + port + "/" + filename)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &models.Camera{Protocol: protocol, Port: port, Ip: ip, FileName: filename, Capturer: cap, Name: name}, nil
}

func (ch *CameraHandler) GetVideoFrame(ctx context.Context, cam *models.Camera) (frame []byte, err error) {
	img := gocv.NewMat()
	defer img.Close()
	for {
		cam.Capturer.Read(&img)
		if img.Empty() {
			continue
		}

		gocv.Resize(img, &img, image.Point{}, float64(0.5), float64(0.5), 0)

		buf, err := gocv.IMEncode(".jpg", img)
		if err != nil {
			ctx.Err()
		}
		return buf.GetBytes(), nil
	}
}
