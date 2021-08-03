package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siconghe/blog/pkg/e"
	"github.com/siconghe/blog/pkg/logging"
	"github.com/siconghe/blog/pkg/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})
			logging.Info(fmt.Sprintf("鉴权失败：%s",e.GetMsg(code)))
			c.Abort()
			return
		}

		c.Next()
	}
}