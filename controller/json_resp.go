package controller

import "github.com/gin-gonic/gin"

func SuccJson(ctx *gin.Context, data interface{}) {
	Resp_json(ctx, 200, 1, "ok", data)
}

func FailJson(ctx *gin.Context, msg string, data interface{}) {
	if len(msg) == 0 || msg == "" {
		msg = "error"
	}

	Resp_json(ctx, 403, 99999, msg, data)
}

func FailMsgJson(ctx *gin.Context, msg string) {
	FailJson(ctx, msg, nil)
}

func FailErrCodeJson(ctx *gin.Context, msg string, errCode int, data interface{}) {
	if len(msg) == 0 || msg == "" {
		msg = "error"
	}

	Resp_json(ctx, 403, errCode, msg, data)
}

func Resp_json(ctx *gin.Context, httpCode int, respCode int, msg string, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"code": respCode,
		"msg":  msg,
		"data": data,
	})
}
