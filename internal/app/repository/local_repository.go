package repository

import (
	"AnalyticsService/internal/app/models"
	"context"
	"errors"
	"fmt"
)

type LocalRepository struct {
	Repos []*models.Camera
}

func (l *LocalRepository) GetListCameras(ctx context.Context) ([]*models.Camera, error) {
	return l.Repos, nil
}

func (l *LocalRepository) GetCameraByPortAndIp(ctx context.Context, port string, ip string) (*models.Camera, error) {
	fmt.Println(len(l.Repos))
	for i := range l.Repos {
		if l.Repos[i].Ip == ip && l.Repos[i].Port == port {
			return l.Repos[i], nil
		}
	}
	return nil, fmt.Errorf("Camera not found")
}

func (l *LocalRepository) DeleteCamera(ctx context.Context, cam *models.Camera) (deleted bool, err error) {
	for i := range l.Repos {
		if l.Repos[i] == cam {
			l.Repos = append(l.Repos[:i], l.Repos[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("camera not found")
}
func (l *LocalRepository) AddCamera(ctx context.Context, cam *models.Camera) (saved bool, err error) {
	l.Repos = append(l.Repos, cam)
	fmt.Println(len(l.Repos))
	return true, nil
}
