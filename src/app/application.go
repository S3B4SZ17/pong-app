package app

import (
	"os"
	"time"

	"github.com/S3B4SZ17/pong-app/src/management"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	router *gin.Engine
)

func StartApp() {

	management.Log.Info("Starting HTTP server ...")
	StartHTTPServer()
}

func StartHTTPServer() {
	gin_mode := os.Getenv("GIN_MODE")

	if gin_mode == "" {
		gin_mode = "release"
		os.Setenv("GIN_MODE", gin_mode)
		gin.SetMode(gin.ReleaseMode)
	}
	httpPort := viper.GetString("http_server.http_port")
	if httpPort == "" {
		httpPort = "8181"
	}

	management.Log.Info("Initializing server with following options: ", zap.String("GIN_MODE", gin_mode), zap.String("HTTP_PORT", httpPort))
	management.Log.Info("Starting application ...")
	router = gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     viper.GetStringSlice("http_server.cors.list_hosts"),
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	mapUrls()

	router.Run(":" + httpPort)
}
