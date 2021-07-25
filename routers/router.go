package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/tom", func(context *gin.Context) {
		context.JSON(200,gin.H{"tom": "喵喵"})
	})
	return r
}
