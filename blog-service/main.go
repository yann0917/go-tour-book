package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yann0917/go-tour-book/blog-service/pkg/tracer"

	"github.com/gin-gonic/gin"
	"github.com/yann0917/go-tour-book/blog-service/global"
	"github.com/yann0917/go-tour-book/blog-service/internal/model"
	"github.com/yann0917/go-tour-book/blog-service/internal/routers"
	"github.com/yann0917/go-tour-book/blog-service/pkg/logger"
	"github.com/yann0917/go-tour-book/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description GO 编程之旅
// @termsOfService http://swagger.io/terms/
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced Shutdown: %v", err)
	}
	log.Println("Server exiting")
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = setupTracer()
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" +
			global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "",
		log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"192.168.1.3:6831",
	)

	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
