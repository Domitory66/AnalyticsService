package main

import (
	gen "AnalyticsService/internal/app/api/camera"
	"AnalyticsService/internal/app/api/video"
	"AnalyticsService/internal/app/handler"
	"AnalyticsService/internal/app/logger"
	"AnalyticsService/internal/app/repository"
	"AnalyticsService/internal/app/server"
	"AnalyticsService/internal/app/service"
	"fmt"
	"net"
	"sync"

	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	//Logger Для общения с SHWService
	log := logger.NewLogger("./logs/grpcLog/", "network.log")

	if err := initConfig(); err != nil {
		log.Log.Fatalf("Error to initialize configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Log.Fatalf("ERROR loading env variables: %s", err.Error())
	}

	log.Log.Info("Starting Analytics Web Cameras service ", slog.String("env", os.Getenv("env")))

	local := &repository.LocalRepository{}
	camHandler := &handler.CameraHandler{}

	repos := repository.New(local)
	handler := handler.New(camHandler)
	services := service.New(handler, repos)

	standLis, err := net.Listen("tcp", viper.GetString("ip")+":"+viper.GetString("standardPort"))
	if err != nil {
		log.Log.Fatalf("Error to listen: %v", err.Error())
	}
	log.Log.Info("Listen port for work with cameras: ", viper.GetString("standardPort"))
	server := server.New(services, log)

	grpcStandardServer := grpc.NewServer()
	gen.RegisterCameraWorkerServer(grpcStandardServer, server)

	videoLis, err := net.Listen("tcp", viper.GetString("ip")+":"+viper.GetString("videoPort"))
	if err != nil {
		log.Log.Fatalf("Error to listen: %v", err.Error())
	}
	log.Log.Info("Listen port for work with video stream: ", viper.GetString("videoPort"))

	grpcVideoStreamServer := grpc.NewServer()
	video.RegisterVideoStreamServer(grpcVideoStreamServer, server)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcStandardServer.Serve(standLis); err != nil {
			log.Log.Fatalf("Failed to serve: %s", err.Error())
		}
	}()
	log.Log.Info("GRPC standard server started on: ", viper.GetString("ip"), ":", viper.GetString("standardPort"))
	fmt.Printf("GRPC standard server started on: %s : %s \n", viper.GetString("ip"), viper.GetString("standardPort"))

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcVideoStreamServer.Serve(videoLis); err != nil {
			log.Log.Fatalf("Failed to serve: %s", err.Error())
		}
	}()
	log.Log.Info("GRPC video server started on: ", viper.GetString("ip"), ":", viper.GetString("videoPort"))
	fmt.Printf("GRPC video server started on: %s : %s \n", viper.GetString("ip"), viper.GetString("videoPort"))
	//TODO  Создать сервер для общения с SmartHomeService
	wg.Wait()

}

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")

	return viper.ReadInConfig()
}
