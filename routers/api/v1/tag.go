package v1

import (
	"GinLearning/models"
	"GinLearning/pkg/setting"
	"GinLearning/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	// 获取 QueryString 请求参数
	name := c.Query("name")
	// 定义俩个 map maps 为查询条件；data 为返回结果
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	// 如果请求参数 name 不为空则放入查询条件
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	util.SuccessResponse(c, data)
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {

}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {

}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {

}
