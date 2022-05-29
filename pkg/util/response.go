package util

import (
	"GinLearning/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessResponse 成功返回
func SuccessResponse(c *gin.Context, data interface{}) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// ErrorResponse 失败返回
func ErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": map[string]interface{}{},
	})
}

func SuccessMessage(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": map[string]interface{}{},
	})
}
