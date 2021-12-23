package config

import (
	"annie-baby/controller"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	video := r.Group("/api/video")
	video.GET("/", controller.VideoList)
	video.GET("/info/:id", controller.VideoInfo)
	video.POST("/download", controller.VideoDownload)
	video.GET("/gen_preview", controller.VideoGenPreview)
	video.DELETE("/delete/:id", controller.VideoDelete)
}
