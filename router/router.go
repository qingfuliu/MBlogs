package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goProject/controller"
	"net/http"
)

func SetUpRouter() *gin.Engine {
	err := controller.RegisteValidator("en")
	if err != nil {
		fmt.Printf("register validation faild,err is %s", err)
	}
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")
	r.Static("/static", "static")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	version := r.Group("/api/v1")
	version.POST("/signup", controller.SingUpHandle)
	//r.GET("//")
	version.POST("/login", controller.LoginHandle)
	version.POST("/communities", controller.QueryBatchCommunityHandle)
	version.POST("/refresh", controller.RefreshToken)

	grantGroup := version.Group("/grant")
	grantGroup.Use(controller.JWTAUTHMiddleWare())
	grantGroup.POST("/communitySignUp", controller.CreateCommunityHandle)
	grantGroup.POST("/PushArticle", controller.PushArticle)

	return r
}
