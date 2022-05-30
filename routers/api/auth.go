package api

import (
	"GinLearning/models"
	"GinLearning/pkg/e"
	"GinLearning/pkg/logging"
	"GinLearning/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// GetAuth 获取token
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		// 此处使用明文账密去做JWT的处理是不合适的
		// 一般会对密码进行非对称加密后存储在库中
		// 而 JWT 的token生成一般也会加盐，并在 Redis 或 MySQL中做维护
		// 但这里只是一个学习的项目 所以不做过多考虑
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	if code == e.SUCCESS {
		util.SuccessResponse(c, data)
	} else {
		util.ErrorResponse(c, code)
	}
}
