package repository

import (
	"AnalyticsService/internal/app/models"
	"context"
)

type MongoRepository struct {
}

// TODO Работа с БД
func (m *MongoRepository) AddCamera(ctx context.Context, userID uint32, camera *models.Camera) (saved bool, err error) {
	return true, nil
}

func (m *MongoRepository) DeleteCamera(ctx context.Context, userID uint32, camera *models.Camera) (deleted bool, err error) {
	return true, nil
}

func (m *MongoRepository) GetListCameras(ctx context.Context, userID uint32) ([]models.Camera, error) {
	return nil, nil
}
