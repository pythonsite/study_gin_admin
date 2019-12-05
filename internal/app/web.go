package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"study_gin_admin/internal/app/config"
	"study_gin_admin/pkg/logger"
	"time"
)

func InitWeb(container *dig.Container) *gin.Engine {
	cfg := config.Global()
	gin.SetMode(cfg.RunMode)

	app := gin.New()


	return app


}

func InitHTTPServer(ctx context.Context, container *dig.Container) func() {
	cfg := config.Global().HTTP
	addr := fmt.Sprintf("%s:%d",cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:              addr,
		Handler:           InitWeb(container),
		ReadTimeout:       5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.Printf(ctx, "HTTP服务开始启动，地址监听在：[%s]", addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Errorf(ctx, err.Error())
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx,time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf(ctx, err.Error())
		}
	}
}