package service

import (
	"AnalyticsService/internal/app/handler"
	"AnalyticsService/internal/app/repository"
)

type Service struct {
	Handler    *handler.Handler
	Repository *repository.Repository
}

func New(h *handler.Handler, r *repository.Repository) *Service {
	return &Service{Handler: h, Repository: r}
}
