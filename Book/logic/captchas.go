package logic

import (
	"Book/model"
	"Book/tools"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func SetEmailCapt() gin.HandlerFunc {
	return func(context *gin.Context) {
		var num Other
		_ = context.ShouldBind(&num)
		//生成验证码
		randnum := rand.Intn(89999) + 10000
		if model.SetEmailCode(num.Email, randnum) != nil {
			context.JSON(http.StatusInternalServerError, tools.EmailParamErr)
		}
		context.JSON(http.StatusOK, tools.OK)
	}
}
func GetEmailCapt() gin.HandlerFunc {
	return func(context *gin.Context) {
		var num Other
		_ = context.ShouldBind(&num)
		//生成验证码
		code, bool2 := model.GetEmailCode(num.Email)

		if !bool2 || code != num.Num {
			context.JSON(http.StatusInternalServerError, tools.EmailErr)
		}
		context.JSON(http.StatusOK, tools.OK)
	}
}
