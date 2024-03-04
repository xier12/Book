package router

import (
	"Book/logic"
	"Book/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	UserIndex := r.Group("")
	UserIndex.Use(middleware.CheckCookie)
	UserIndex.POST("/regis", logic.RegisUser())
	UserIndex.POST("/loginByName", logic.Login())
	UserIndex.POST("/loginByEmail", logic.LoginByEmail())
	UserIndex.POST("/SendEmail", logic.SendEmail())
	UserIndex.POST("/loginByTel", logic.LoginByTel())
	UserIndex.POST("/SendTel", logic.SendTel())
	UserIndex.POST("/loginByNameAndPRC", logic.LoginByNameAndPRC())

	UserIndex.GET("/captcha", logic.Captcha())
	UserIndex.POST("/captcha/verify", logic.CaptchaVerify())

	DoBook := UserIndex.Group("")
	DoBook.Use(middleware.JWTAuth())
	//DoBook.Use(middleware.CheckCookie)
	DoBook.POST("/UpLoadUserPhoto", logic.UpLoadUserPhoto())
	UserIndex.GET("/UserInfo", logic.UserInfo())
	DoBook.GET("/MyRecord", logic.BorrowRecordToUser())
	DoBook.GET("/RootRecord", logic.BorrowRecordToRoot())
	DoBook.PUT("/ReturningBook", logic.ReturningBook())
	DoBook.POST("/BorrowBook", logic.BorrowBook())
	DoBook.POST("/BuyBook", logic.PayUrl())
}
