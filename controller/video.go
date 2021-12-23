package controller

import "github.com/gin-gonic/gin"

// get
func VideoList(c *gin.Context) {
	query := c.DefaultQuery("ass", "kick")
	c.JSON(200, gin.H{
		"query": query,
	})
}

//get
func VideoInfo(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
	})

}

type downloadRequest struct {
	ID     int64  `form:"id"`
	Stream string `form:"stream"`
}

//post
func VideoDownload(c *gin.Context) {
	form := downloadRequest{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(403, gin.H{
			"msg": "fuck up",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "ok",
		"form": form,
	})
}

//get
func VideoGenPreview(c *gin.Context) {

}

//delete
func VideoDelete(c *gin.Context) {

}
