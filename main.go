package main

import (
	"annie-baby/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//设置路由
	config.SetRoutes(r)

	r.Run(":3000")
}
