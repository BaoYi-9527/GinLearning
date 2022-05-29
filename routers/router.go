package routers

import (
	"GinLearning/pkg/setting"
	v1 "GinLearning/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	// 创建不同的 HTTP 方法绑定到 Handlers 中
	r.GET("/test", func(context *gin.Context) {
		// gin.H{...} 就是一个 map[string]interface{}
		// gin.Context 是 gin 中的上下文 其允许在中间件之间传递变量、管理流、验证JSON请求、响应JSON请求等
		context.JSON(200, gin.H{
			"message": "test",
		})
	})

	// 注册api路由
	apiV1 := r.Group("/api/v1")
	{
		// 获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		// 新建标签
		apiV1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
