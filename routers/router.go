package routers

import (
	_ "GinLearning/docs"
	"GinLearning/middleware/jwt"
	"GinLearning/pkg/setting"
	"GinLearning/routers/api"
	v1 "GinLearning/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// 获取鉴权 Token
	r.GET("/auth", api.GetAuth)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册api路由
	apiV1 := r.Group("/api/v1")
	// 接入鉴权中间件
	apiV1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		// 新建标签
		apiV1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
		// 获取文档详情
		apiV1.GET("/article/:id", v1.GetArticle)
		// 获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		// 新建文章
		apiV1.POST("/articles", v1.AddArticle)
		// 更新指定文章
		apiV1.PUT("/articles/:id", v1.EditArticle)
		// 删除指定文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
