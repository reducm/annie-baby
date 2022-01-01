package controller

import (
	"annie-baby/ent"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
)

type VideoController struct {
	DB *ent.Client
}

// get
func (c *VideoController) List(ctx *gin.Context) {
	query := ctx.DefaultQuery("ass", "kick")

	videos, err := c.DB.Video.Query().All(context.Background())

	if err != nil {
		pretty.Logf("bind json err %v", err)
	}

	ctx.JSON(200, gin.H{
		"query":  query,
		"videos": videos,
	})
}

//POST 只拉info
func (c *VideoController) Info(ctx *gin.Context) {
	var json VideoInfoRO
	err := ctx.ShouldBindJSON(&json)

	if err != nil {
		pretty.Logf("bind json err %v", err)
		FailJson(ctx, err.Error(), err)
		return
	}

	SuccJson(ctx, json)
}

//GET 获取详情
func (c *VideoController) Show(ctx *gin.Context) {
}

type downloadRequest struct {
	ID     int64  `form:"id"`
	Stream string `form:"stream"`
}

//POST 下载,
// 有ID的时候1.查出 2.有无指定steam再调用默认下载
// 没有ID的时候直接使用url创建model再下载
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

// 默认下载完成会生成略缩图, 这个用于生成hls
func (c *VideoController) GenPreview(ctx *gin.Context) {

}

//delete
func (c *VideoController) Delete(ctx *gin.Context) {

}
