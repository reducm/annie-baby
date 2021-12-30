package config

import (
	"annie-baby/controller"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {

	video_controller := &controller.VideoController{ApplicationController: controller.ApplicationController{DB: DBClient, Name: "video"}}

	video := r.Group("/api/video")
	video.GET("/", video_controller.List)
	video.GET("/info/:id", video_controller.Info)
	video.POST("/download", video_controller.Delete)
	video.GET("/gen_preview", video_controller.GenPreview)
	video.DELETE("/delete/:id", video_controller.Delete)
}
