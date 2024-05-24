package server

import (
	gen "AnalyticsService/internal/app/api/camera"
	"AnalyticsService/internal/app/api/video"
	"AnalyticsService/internal/app/logger"
	"AnalyticsService/internal/app/service"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	gen.UnimplementedCameraWorkerServer
	video.UnimplementedVideoStreamServer
	service.Service
	logger.Logger
}

func New(s *service.Service, l *logger.Logger) *serverAPI {
	return &serverAPI{Service: *s, Logger: *l}
}

// TODO валидация при входе в микросервис данных
// TODO встраивание лога

func (s *serverAPI) FindCamera(ctx context.Context, in *gen.FindCameraRequest) (*gen.FindCameraResponse, error) {
	_, err := s.Service.Handler.FindCamera(ctx, in.Camera.Name, in.Camera.Port, in.Camera.Ip, in.Camera.Protocol, in.Camera.Filename)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Camera not found")
	}
	return &gen.FindCameraResponse{Found: true}, nil
}

func (s *serverAPI) AddCamera(ctx context.Context, in *gen.AddCameraRequest) (*gen.AddCameraResponse, error) {
	s.Logger.Log.Println("method AddCamera: ", in.GetCamera())
	camera, err := s.Service.Handler.FindCamera(ctx, in.Camera.Name, in.Camera.Port, in.Camera.Ip, in.Camera.Protocol, in.Camera.Filename)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Camera not found")
	}

	saved, err := s.Service.Repository.AddCamera(ctx, camera)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Camera already exist")
	}

	return &gen.AddCameraResponse{Saved: saved}, nil
}

func (s *serverAPI) DeleteCamera(ctx context.Context, in *gen.DeleteCameraRequest) (*gen.DeleteCameraResponse, error) {
	camera, err := s.Repository.GetCameraByPortAndIp(ctx, in.Camera.Port, in.Camera.Ip)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Camera not found")
	}

	deleted, err := s.Service.Repository.DeleteCamera(ctx, camera)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Camera not deleted")
	}
	return &gen.DeleteCameraResponse{Deleted: deleted}, nil
}

func (s *serverAPI) GetAllCameras(ctx context.Context, in *gen.GetAllCamerasRequest) (*gen.GetAllCamerasResponse, error) {
	list, err := s.Repository.CameraRepository.GetListCameras(ctx)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	genCam := make([]*gen.Camera, len(list))
	for i, cam := range list {
		genCam[i] = &gen.Camera{Name: cam.Name, Port: cam.Port, Ip: cam.Ip, Protocol: cam.Protocol, Filename: cam.FileName}
	}

	return &gen.GetAllCamerasResponse{Cameras: genCam}, nil
}

func (s *serverAPI) GetCameraByPortAndIp(ctx context.Context, in *gen.GetCameraRequest) (*gen.GetCameraResponse, error) {
	cam, err := s.Repository.GetCameraByPortAndIp(ctx, in.Port, in.Ip)
	fmt.Print(cam)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &gen.GetCameraResponse{Camera: &gen.Camera{Name: cam.Name, Port: cam.Port, Ip: cam.Ip, Protocol: cam.Protocol, Filename: cam.FileName}}, nil
}

func (s *serverAPI) GetVideoFromCamera(ctx context.Context, in *video.ImageRequest) (*video.ImageResponse, error) {
	cam, err := s.Repository.GetCameraByPortAndIp(ctx, in.Port, in.Ip)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	frame, err := s.Service.Handler.GetVideoFrame(ctx, cam)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	frame, _, err = s.Service.Handler.FindGesture(ctx, frame, cam)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &video.ImageResponse{Ip: in.Ip, Port: in.Port, Image: frame}, nil
}

func (s *serverAPI) StopVideoStream(ctx context.Context, in *video.StopRequest) (*video.StopResponse, error) {
	// res, err := s.Service.Handler.StopVideoStream(ctx, in.Ip, in.Port)
	// if err != nil {
	// 	return nil, status.Error(codes.NotFound, err.Error())
	// }
	return &video.StopResponse{Stopped: true}, nil
}
