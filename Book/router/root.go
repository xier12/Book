package router

import (
	"Book/logic"
	"github.com/gin-gonic/gin"
)

func InitRootRouter(r *gin.Engine) {
	RootIndex := r.Group("")
	//RootIndex.Use(middleware.CheckCookie)
	//增
	RootIndex.POST("/Add", logic.AddBook())
	RootIndex.POST("/UpLoadBookPhoto", logic.UpLoadBookPhoto())
	//删
	RootIndex.DELETE("/Delete", logic.DeleteBook())
	//改
	RootIndex.PUT("/Update", logic.UpdateBook())
	//查
	RootIndex.GET("/BookIndex", logic.FindAllBook())
	RootIndex.GET("/BookInfo", logic.FindBookByName())
}

// 显示图片
