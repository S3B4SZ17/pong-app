package app

import (
	"github.com/S3B4SZ17/pong-app/src/controllers"
)

func mapUrls() {

	public := router.Group("/api")
	public.GET("/ping", controllers.Ping)

}
