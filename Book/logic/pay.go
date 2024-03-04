package logic

import (
	"Book/model"
	"Book/tools"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func PayUrl() gin.HandlerFunc {
	return func(context *gin.Context) {
		n, _ := context.Get("userPower")
		power := n.(string)
		myrole := tools.FindRole(power)
		aaaaaa, _ := myrole.Can("watch", "/BuyBook")
		if !aaaaaa {
			context.JSON(http.StatusOK, tools.Auth)
		}
		var book model.Book
		name := context.Query("name")
		num := context.Query("num")
		book = model.FindBookByName(name)
		orderID0, _ := tools.Snow.GenerateID()
		orderID := strconv.FormatInt(orderID0, 10)
		a, _ := context.Get("userID")
		aa := a.(int64)
		b, _ := strconv.ParseInt(num, 10, 64)
		order := model.Order{
			Bid:        book.Id,
			Uid:        aa,
			Num:        b,
			Price:      float32(b) * book.Price,
			Createtime: time.Now().Format("2006-01-02 15:04:05"),
		}
		if err1 := model.AddOrder(order); err1 != nil {
			context.JSON(http.StatusOK, tools.ECode{
				Code:    11000,
				Message: "系统错误",
			})
		}
		url, err := model.AliPayClient.Pay(tools.Order{
			ID:          orderID,
			Subject:     "ttms购票:" + orderID,
			TotalAmount: book.Price,
			Code:        tools.LaptopWebPay,
		})
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusOK, tools.ECode{
				Code:    11000,
				Message: "系统错误",
			})
			return
		}
		context.Redirect(http.StatusTemporaryRedirect, url)
	}
}

// 重定向到支付宝二维码
func payUrl(c *gin.Context) {
	orderID0, _ := tools.Snow.GenerateID()
	orderID := strconv.FormatInt(orderID0, 10)
	url, err := model.AliPayClient.Pay(tools.Order{
		ID:          orderID,
		Subject:     "ttms购票:" + orderID,
		TotalAmount: 30,
		Code:        tools.LaptopWebPay,
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "系统错误")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// 支付后页面的重定向界面
func Callback(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := model.AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "校验失败")
		return
	}
	c.JSON(http.StatusOK, "支付成功:"+orderID)
}

// 支付成功后支付宝异步通知
func Notify(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := model.AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("支付成功:" + orderID)
	// 做自己的事
}
