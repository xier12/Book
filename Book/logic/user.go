package logic

import (
	"Book/model"
	"Book/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Name         string `json:"name" form:"name"`
	Password     string `json:"password" form:"password"`
	CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}
type DateAndPhoto struct {
	User   model.User
	UpPath []string
}
type Code struct {
	Key  string `json:"key" form:"key"`
	Code int    `json:"code" form:"code"`
}

// 注册用户
func RegisUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var user model.User
		err := context.ShouldBind(&user)
		fmt.Println(user)
		user.Uuid, _ = tools.Snow.GenerateID()
		user.Ctime = time.Now().Format("2006-01-02 15:04:05")
		user.UTime = time.Now().Format("2006-01-02 15:04:05")
		user.Power = "user"
		user.Password = tools.EncryptV1(user.Password)
		if err == nil {
			a := model.AddUser(user)
			fmt.Println(a)
			context.JSON(http.StatusOK, tools.OK)
		}

	}
}

// 登录
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		n, _ := c.Get("userPower")
		power := n.(string)
		myrole := tools.FindRole(power)
		aaaaaa, _ := myrole.Can("watch", "/loginByName")
		if !aaaaaa {
			c.JSON(http.StatusOK, tools.Auth)
		}
		var user User
		var user1 model.User
		//获取前台数据
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusOK, tools.ECode{

				Message: err.Error(), //这里有风险
			})
		}
		logrus.Println(user)

		param1 := tools.CaptchaData{
			CaptchaId: user.CaptchaId,
			Data:      user.CaptchaValue,
		}
		if !tools.CaptchaVerify(param1) {
			c.JSON(http.StatusOK, tools.ECode{
				Code:    10008,
				Message: "验证失败",
			})
		}

		user1 = model.FindUserByName(user.Name)
		if tools.EncryptV1(user.Password) == user1.Password && user.Password != "" && user1.Password != "" {
			fmt.Println(1)
			//使用session代替cookie保持登录态
			_ = model.SetSession(c, user.Name, user1.Id)
			a, _ := model.GenToken(user1.Id, user1.Name, user1.Power)
			b, _ := model.ParseToken(a)
			// 将当前请求的userID信息保存到请求的上下文c上
			c.Set("userID", b.UserId)
			c.Set("userName", b.Username)
			c.Set("userPower", b.UserPower)
			//向前端输出数据
			c.JSON(http.StatusOK, tools.OK)
			//c.JSON(http.StatusOK, tools.ECode{
			//	//code赋值为0,使用前端重定向
			//	Message: "登录成功111",
			//})
		}

		c.JSON(http.StatusOK, tools.UserParamErr)
	}
}
func LoginByNameAndPRC() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		//获取前台数据
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusOK, tools.ECode{
				Message: err.Error(), //这里有风险
			})
		}
		reslut := tools.PRC(user.Name, user.Password, user.CaptchaId, user.CaptchaValue)

		//使用session代替cookie保持登录态
		_ = model.SetSession(c, user.Name, reslut.Id)
		a, _ := model.GenToken(reslut.Id, reslut.Name, reslut.Power)
		fmt.Println(a)
		c.Header("Authorization", a)
		b, _ := model.ParseToken(a)
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userID", b.UserId)
		c.Set("userName", b.Username)
		c.Set("userPower", b.UserPower)
		//向前端输出数据
		c.JSON(http.StatusOK, tools.OK)
		//c.JSON(http.StatusOK, tools.ECode{
		//	//code赋值为0,使用前端重定向
		//	Message: "登录成功111",
		//})
	}
}
func SendEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code Code
		//获取前台数据
		if err := c.ShouldBind(&code); err != nil {
			c.JSON(http.StatusOK, tools.ECode{
				Message: err.Error(), //这里有风险
			})
		}
		fmt.Println(code)
		randnum := rand.Intn(89999) + 10000
		fmt.Println(randnum)
		model.SetEmailCode(code.Key, randnum)
		tools.SendEmail(randnum, code.Key)
		c.JSON(http.StatusOK, tools.OK)
	}
}
func SendTel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code Code
		//获取前台数据
		if err := c.ShouldBind(&code); err != nil {
			c.JSON(http.StatusOK, tools.ECode{
				Message: err.Error(), //这里有风险
			})
		}
		randnum := rand.Intn(89999) + 10000
		fmt.Println(randnum)
		model.SendMessage(code.Key, strconv.Itoa(randnum))
		c.JSON(http.StatusOK, tools.OK)
	}
}
func LoginByEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code Code
		var user model.User
		//获取前台数据
		if err := c.ShouldBind(&code); err != nil {
			c.JSON(http.StatusOK, tools.ECode{
				Message: err.Error(), //这里有风险
			})
		}
		fmt.Println(1)
		if a, _ := model.GetEmailCode(code.Key); a == code.Code {
			user = model.FindUserByEmail(code.Key)
			//使用session代替cookie保持登录态
			_ = model.SetSession(c, user.Name, user.Id)
			a1, _ := model.GenToken(user.Id, user.Name, user.Power)
			b, _ := model.ParseToken(a1)

			// 将当前请求的userID信息保存到请求的上下文c上
			c.Set("userID", b.UserId)
			c.Set("userName", b.Username)
			c.Set("userPower", b.UserPower)
			c.JSON(http.StatusOK, tools.OK)
		} else {
			c.JSON(http.StatusOK, tools.ECode{
				Code:    10008,
				Message: "验证失败",
			})
		}
	}
}
func LoginByTel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code Code
		var user model.User
		//获取前台数据
		if err := c.ShouldBind(&code); err != nil {
			c.JSON(http.StatusOK, tools.ECode{
				Message: err.Error(), //这里有风险
			})
		}
		if a := model.GetMessage(code.Key); a == strconv.Itoa(code.Code) {
			user = model.FindUserByTel(code.Key)
			//使用session代替cookie保持登录态
			_ = model.SetSession(c, user.Name, user.Id)
			a1, _ := model.GenToken(user.Id, user.Name, user.Power)
			c.Header("Authorization", a1)
			b, _ := model.ParseToken(a1)
			// 将当前请求的userID信息保存到请求的上下文c上
			c.Set("userID", b.UserId)
			c.Set("userName", b.Username)
			c.Set("userPower", b.UserPower)
			c.JSON(http.StatusOK, tools.OK)
		} else {
			c.JSON(http.StatusOK, tools.ECode{
				Code:    10008,
				Message: "验证失败",
			})
		}

	}
}
func LoginByTel1() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code Code
		var user model.User
		//获取前台数据
		if err := c.ShouldBind(&code); err != nil {
			c.JSON(http.StatusOK, tools.ECode{
				Message: err.Error(), //这里有风险
			})
		}
		if a := model.GetMessage(code.Key); a == strconv.Itoa(code.Code) {
			user = model.FindUserByTel(code.Key)
			//使用session代替cookie保持登录态
			_ = model.SetSession(c, user.Name, user.Id)

			a1, _ := model.GenToken(user.Id, user.Name, user.Power)
			b, _ := model.ParseToken(a1)
			// 将当前请求的userID信息保存到请求的上下文c上
			c.Set("userID", b.UserId)
			c.Set("userName", b.Username)
			c.Set("userPower", b.UserPower)
			c.JSON(http.StatusOK, tools.OK)
		}
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10008,
			Message: "验证失败",
		})
	}
}

// 显示用户信息
func UserInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		mc, _ := model.ParseToken(authHeader)
		//n, _ := context.Get("userPower")
		//power := n.(string)
		myrole := tools.FindRole(mc.UserPower)
		aaaaaa, _ := myrole.Can("watch", "/UserInfo")
		if !aaaaaa {
			context.JSON(http.StatusOK, tools.Auth)
		}
		//jname, _ := context.Get("userName")
		//name := jname.(string)
		var user model.User
		var uppath []string
		if user = model.FindUserByName(mc.Username); user.Id <= 0 {
			context.JSON(http.StatusOK, tools.ParamErr)
		}

		fmt.Printf("user:%s\n", user)

		if uppath = model.FindUserPhotoByUserId(user.Id); len(uppath) > 0 {
			fmt.Printf("user:%s\n", user)
			context.JSON(http.StatusOK, tools.ECode{
				Code: 12,
				Data: DateAndPhoto{
					User:   user,
					UpPath: uppath,
				},
			})
		}
	}
}

// 图片验证码
func Captcha() gin.HandlerFunc {
	return func(context *gin.Context) {
		n, _ := context.Get("userPower")
		power := n.(string)
		myrole := tools.FindRole(power)
		aaaaaa, _ := myrole.Can("watch", "/captcha")
		if !aaaaaa {
			context.JSON(http.StatusOK, tools.Auth)
		}
		captcha, err := tools.CaptchaGenerate()
		if err != nil {
			context.JSON(http.StatusOK, tools.ECode{
				Code:    10005,
				Message: err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, tools.ECode{
			Data: captcha,
		})
	}
}
func CaptchaVerify() gin.HandlerFunc {
	return func(context *gin.Context) {
		n, _ := context.Get("userPower")
		power := n.(string)
		myrole := tools.FindRole(power)
		aaaaaa, _ := myrole.Can("watch", "/captcha/verify")
		if !aaaaaa {
			context.JSON(http.StatusOK, tools.Auth)
		}
		var param tools.CaptchaData
		if err := context.ShouldBind(&param); err != nil {
			context.JSON(http.StatusOK, tools.ParamErr)
			return
		}

		fmt.Printf("参数为：%+v", param)
		if !tools.CaptchaVerify(param) {
			context.JSON(http.StatusOK, tools.ECode{
				Code:    10008,
				Message: "验证失败",
			})
			return
		}
		context.JSON(http.StatusOK, tools.OK)
	}
}
