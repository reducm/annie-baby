package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
)

type VideoController struct {
	ApplicationController
}

// get
func (c *VideoController) List(ctx *gin.Context) {
	query := ctx.DefaultQuery("ass", "kick")

	videos, err := c.ApplicationController.DB.Video.Query().All(context.Background())

	if err != nil {
		pretty.Logf(" err %v", err)
	}

	ctx.JSON(200, gin.H{
		"query":  query,
		"videos": videos,
	})
}

//get
func (c *VideoController) Info(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, gin.H{
		"id": id,
	})

}

type downloadRequest struct {
	ID     int64  `form:"id"`
	Stream string `form:"stream"`
}

//post
func (c *VideoController) Download(ctx *gin.Context) {
	form := downloadRequest{}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(403, gin.H{
			"msg": "fuck up",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"msg":  "ok",
		"form": form,
	})
}

//get
func (c *VideoController) GenPreview(ctx *gin.Context) {

}

//delete
func (c *VideoController) Delete(ctx *gin.Context) {

}
