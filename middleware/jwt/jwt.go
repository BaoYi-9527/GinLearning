package jwt

import (
	"GinLearning/pkg/e"
	"GinLearning/pkg/util"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int

		code = e.SUCCESS
		authorization := context.GetHeader("Authorization")
		token := strings.TrimPrefix(authorization, "Bearer ")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				//
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				//
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		//
		if code != e.SUCCESS {
			util.ErrorResponse(context, code)
			context.Abort()
			return
		}
		context.Next()
	}
}
