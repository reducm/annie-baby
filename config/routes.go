package config

import (
	"annie-baby/controller"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {

	video_controller := &controller.VideoController{DB: DBClient}

	video := r.Group("/api/video")
	video.GET("/", video_controller.List)
	video.GET("/:id", video_controller.Show)
	video.POST("/download", video_controller.Download)
	video.POST("/info", video_controller.Info)
	video.GET("/gen_preview", video_controller.GenPreview)
	video.DELETE("/delete/:id", video_controller.Delete)
}
