package handler

import (
	"AnalyticsService/internal/app/models"
	"context"
)

type VideoWorker interface {
	FindGesture(
		ctx context.Context,
		camera *models.Camera,
	) (frame []byte, typeGesture string, err error)
}

type CameraWorker interface {
	FindCamera(
		ctx context.Context,
		name string,
		port string,
		ip string,
		protocol string,
		filename string,
	) (camera *models.Camera, err error)
	GetVideoFrame(
		ctx context.Context,
		camera *models.Camera,
	) (frame []byte, err error)
}

type Handler struct {
	CameraWorker
	VideoWorker
}

func New(cw CameraWorker) *Handler {
	return &Handler{CameraWorker: cw}
}
