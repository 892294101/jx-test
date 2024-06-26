// 启动程序
// author xiaoRui

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jx/jxserver/common/config"
	_ "github.com/jx/jxserver/docs"
	"github.com/jx/jxserver/pkg/db"
	"github.com/jx/jxserver/pkg/log"
	"github.com/jx/jxserver/pkg/redis"
	"github.com/jx/jxserver/router"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title 通用后台管理系统
// @version 1.0
// @description 后台管理系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log := log.Log()
	gin.SetMode(config.Config.Server.Model)
	router := router.InitRouter()
	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: router,
	}
	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info("listen: %s \n", err)
		}
		log.Info("listen: %s \n", config.Config.Server.Address)
	}()
	quit := make(chan os.Signal)
	//监听消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

// 初始化连接
func init() {
	// mysql
	db.SetupDBLink()
	// redis
	redis.SetupRedisDb()
}
