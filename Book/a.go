package main

import (
	"Book/tools"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

const (
	kAppID               = "2021000121601691"
	kPrivateKey          = "XXXX"
	kServerDomain        = "http://XXXX:7999"
	AppPublicCertPath    = "cert/appCertPublicKey.crt"         // app公钥证书路径
	AliPayRootCertPath   = "cert/alipayRootCert.crt"           // alipay根证书路径
	AliPayPublicCertPath = "cert/alipayCertPublicKey_RSA2.crt" // alipay公钥证书路径
	NotifyURL            = kServerDomain + "/notify"
	ReturnURL            = kServerDomain + "/callback"
	IsProduction         = false
)

var AliPayClient *tools.AliPayClient

func main() {
	var s = gin.Default()
	s.GET("/alipay", payUrl)
	s.GET("/callback", callback)
	s.POST("/notify", notify)
	s.Run(":8080")
}

// 重定向到支付宝二维码
func payUrl(c *gin.Context) {
	orderID := "1"
	url, err := AliPayClient.Pay(tools.Order{
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
func callback(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "校验失败")
		return
	}
	c.JSON(http.StatusOK, "支付成功:"+orderID)
}

// 支付成功后支付宝异步通知
func notify(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("支付成功:" + orderID)
	// 做自己的事
}
