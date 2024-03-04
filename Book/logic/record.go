package logic

import (
	"Book/model"
	"Book/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Other struct {
	Num   int    `json:"num" form:"num"`
	Email string `json:"email" form:"email"`
}

func BorrowBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		n, _ := context.Get("userPower")
		power := n.(string)
		myrole := tools.FindRole(power)
		aaaaaa, _ := myrole.Can("watch", "/BorrowBook")
		if !aaaaaa {
			context.JSON(http.StatusOK, tools.Auth)
		}
		var num Other
		_ = context.ShouldBind(&num)
		var re model.BookWithUser
		name := context.Query("name")
		b := model.FindBookByName(name)
		if (b.Num - b.DNum) < num.Num {
			context.JSON(http.StatusBadRequest, tools.ParamErr)
		}
		//a为从jwt中取出的1
		//a := int64(2)
		jid, _ := context.Get("userID")
		id := jid.(int64)
		re.Uid = id
		re.Bid = b.Id
		re.CTime = time.Now().Format("2006-01-02 15:04:05")
		re.UTime = re.CTime
		re.Num = num.Num
		model.BorrowBooks(re)
		context.JSON(http.StatusOK, tools.OK)
	}
}
func ReturningBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		n, _ := context.Get("userPower")
		power := n.(string)
		myrole := tools.FindRole(power)
		aaaaaa, _ := myrole.Can("watch", "/ReturningBook")
		if !aaaaaa {
			context.JSON(http.StatusOK, tools.Auth)
		}
		var num Other
		_ = context.ShouldBind(&num)
		jid, _ := context.Get("userID")
		id := jid.(int64)
		name := context.Query("name")
		b := model.FindBookByName(name).Id
		model.ReturnBook(id, b, num.Num)
		context.JSON(http.StatusOK, tools.OK)
	}
}
func BorrowRecordToRoot() gin.HandlerFunc {
	return func(context *gin.Context) {
		n, _ := context.Get("userPower")
		power := n.(string)
		myrole := tools.FindRole(power)
		aaaaaa, _ := myrole.Can("watch", "/RootRecord")
		if !aaaaaa {
			context.JSON(http.StatusOK, tools.Auth)
		}
		var recordlist []model.BookWithUser
		recordlist = model.FindAllRecord()
		context.JSON(http.StatusOK, tools.ECode{Data: recordlist})
	}
}
func BorrowRecordToUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		n, _ := context.Get("userPower")
		power := n.(string)
		myrole := tools.FindRole(power)
		aaaaaa, _ := myrole.Can("watch", "/MyRecord")
		if !aaaaaa {
			context.JSON(http.StatusOK, tools.Auth)
		}
		var book []model.BookWithUser
		jid, _ := context.Get("userID")
		id := jid.(int64)
		book = model.FindRecordByUserId(id)
		context.JSON(http.StatusOK, tools.ECode{Data: book})
	}
}
