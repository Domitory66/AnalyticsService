package repository

import (
	"AnalyticsService/internal/app/models"
	"context"
)

type CameraRepository interface {
	AddCamera(
		ctx context.Context,
		camera *models.Camera,
	) (saved bool, err error)
	DeleteCamera(
		ctx context.Context,
		camera *models.Camera,
	) (deleted bool, err error)
	GetListCameras(
		ctx context.Context,
	) ([]*models.Camera, error)
	GetCameraByPortAndIp(
		ctx context.Context,
		port string,
		ip string,
	) (*models.Camera, error)
}

type Repository struct {
	CameraRepository
}

func New(cr CameraRepository) *Repository {
	return &Repository{CameraRepository: cr}
}
