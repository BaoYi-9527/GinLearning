package v1

import (
	"GinLearning/models"
	"GinLearning/pkg/e"
	"GinLearning/pkg/setting"
	"GinLearning/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetTags
// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} models.Tag
// @Failure 403 {object} string {"code":400,"data":[],"msg":"请求参数错误"}
// @Router /api/v1/tags [get]
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

// AddTag
// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	// 验证
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为 100 字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为 100 字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistsTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
			util.SuccessMessage(c)
		} else {
			code = e.ERROR_EXIST_TAG
			util.ErrorResponse(c, code)
		}
	}

}

// EditTag
// @Summary 修改文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许 0 或 1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为 100 字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为 100 字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistsTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
			util.SuccessMessage(c)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
			util.ErrorResponse(c, code)
		}
	}

}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistsTagById(id) {
			models.DeleteTag(id)
			util.SuccessMessage(c)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
			util.ErrorResponse(c, code)
		}
	}
}
