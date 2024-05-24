package main

import (
	gen "AnalyticsService/internal/app/api/camera"
	"AnalyticsService/internal/app/api/image/prot"
	"AnalyticsService/internal/app/api/video"
	"AnalyticsService/internal/app/handler"
	"AnalyticsService/internal/app/logger"
	"AnalyticsService/internal/app/repository"
	"AnalyticsService/internal/app/server"
	"AnalyticsService/internal/app/service"
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()
	conProcessor, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Log.Fatal("Analytics service not started ", err.Error())
	}
	clientGestureRecognition := prot.NewImageWorkerClient(conProcessor)
	videoHandler := &handler.VideoHandler{clientGestureRecognition}

	repos := repository.New(local)
	handler := handler.New(camHandler, videoHandler)
	camService := service.New(handler, repos)

	standLis, err := net.Listen("tcp", viper.GetString("ip")+":"+viper.GetString("standardPort"))
	if err != nil {
		log.Log.Fatalf("Error to listen: %v", err.Error())
	}
	log.Log.Info("Listen port for work with cameras: ", viper.GetString("standardPort"))
	server := server.New(camService, log)

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
	wg.Wait()

}

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")

	return viper.ReadInConfig()
}
