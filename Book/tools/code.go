package tools

import "fmt"

var (
	OK   = ECode{Code: 0, Message: "成功"}
	Auth = ECode{Code: 10000, Message: "权限不足"}

	UserNotLogin = ECode{Code: 11001, Message: "用户未登录"}
	UserParamErr = ECode{Code: 11002, Message: "用户参数错误"}
	UserRegisErr = ECode{Code: 11003, Message: "用户注册参数错误"}
	UserFailed   = ECode{Code: 11004, Message: "用户创建失败"}

	BookFailed = ECode{Code: 12001, Message: "图书不存在"}
	BookErr    = ECode{Code: 12002, Message: "图书操作出现错误"}
	BookEmpty  = ECode{Code: 12003, Message: "图书不存在"}

	ParamErr = ECode{Code: 13001, Message: "参数错误"}

	EmailParamErr = ECode{Code: 14001, Message: "邮件验证码生成错误"}
	EmailErr      = ECode{Code: 14002, Message: "邮件验证码错误"}

	PhotoUpLoadErr = ECode{Code: 15002, Message: "图片上传错误"}
)

type ECode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (e *ECode) String() string {
	return fmt.Sprintf("code:%d,message:%s", e.Code, e.Message)
}
