package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Ping(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": viper.GetString("pong.message"),
	})
}
