package routes

import (
	"github.com/georgehao/wechat/app/middleware"
	"github.com/georgehao/wechat/providers"
	"github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.PUT("/login", providers.LoginController.Login)

	group := r.Group("")
	group.Use(middleware.AuthMiddlerware.MiddlewareFunc())
	{
		group.GET("picture_books/isbns/:isbn", providers.ScanCodeController.Scan)
		group.GET("/picture_books/records", providers.ScanCodeController.Index)
		group.POST("/picture_books/records", providers.ScanCodeController.Store)
	}

	return r
}
