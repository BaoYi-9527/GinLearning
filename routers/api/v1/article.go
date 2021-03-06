package v1

import (
	"GinLearning/models"
	"GinLearning/pkg/e"
	"GinLearning/pkg/logging"
	"GinLearning/pkg/setting"
	"GinLearning/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetArticle 获取文章详情
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID 必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistsArticleByID(id) {
			data = models.GetArticle(id)
			util.SuccessResponse(c, data)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
			util.ErrorResponse(c, code)
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
}

// GetArticles 文章列表
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	if !valid.HasErrors() {
		data["list"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
		util.SuccessResponse(c, data)
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistsTagById(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["content"] = content
			data["desc"] = desc
			data["state"] = state
			data["created_by"] = createdBy
			models.AddArticle(data)
			util.SuccessMessage(c)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
			util.ErrorResponse(c, code)
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为 100 字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为 255 字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为 65535 字符")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为 100 字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistsArticleByID(id) {
			if models.ExistsTagById(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				data["modified_by"] = modifiedBy
				models.EditArticle(id, data)
				util.SuccessMessage(c)
			} else {
				code = e.ERROR_NOT_EXIST_TAG
				util.ErrorResponse(c, code)
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
			util.ErrorResponse(c, code)
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistsArticleByID(id) {
			models.DeleteArticle(id)
			util.SuccessMessage(c)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
			util.ErrorResponse(c, code)
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
}
