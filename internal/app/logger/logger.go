package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// type TypeLogger int

// const (
// 	GRPCLogger TypeLogger = iota
// )

type Logger struct {
	Log *logrus.Logger
}

func NewLogger(path string, filename string) *Logger {
	logFile, err := os.OpenFile(path+filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("ERROR opening log file: %v", err)
	}
	log := &logrus.Logger{
		Out:       logFile,
		Formatter: &logrus.TextFormatter{},
		Level:     logrus.DebugLevel,
		ExitFunc:  os.Exit,
	}
	logger := &Logger{Log: log}
	return logger
}

// func (l *Logger) setGRPCLog() {

// }

// func (l *Logger) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
// 	l.Log.Println("UnaryServerInterceptor PRE", info.FullMethod)

// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		l.Log.Print()
// 	}
// }

// func LevelDetect(code codes.Code) logrus.Level {
// 	if code == codes.OK {
// 		return logrus.InfoLevel
// 	}
// 	return logrus.ErrorLevel
// }
