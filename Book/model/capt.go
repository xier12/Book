package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"strconv"
	"time"
)

func SetEmailCode(targetEmail string, code int) error {
	ret := ConnectRedis.Get(context.TODO(), fmt.Sprintf("Code:Emailc:%s", targetEmail))
	if ret == nil {
		return fmt.Errorf("验证码已发送,稍后在试")
	}
	_ = ConnectRedis.Set(context.TODO(), fmt.Sprintf("Code:Email:%s", targetEmail), code, 600*time.Second)
	_ = ConnectRedis.Set(context.TODO(), fmt.Sprintf("Code:Emailc:%s", targetEmail), 100, 60*time.Second)
	return nil
}
func GetEmailCode(targetEmail string) (code int, bool2 bool) {
	ret, err := ConnectRedis.Get(context.TODO(), fmt.Sprintf("Code:Email:%s", targetEmail)).Result()
	if err != nil {
		code = 0
		bool2 = false
		return code, bool2
	} else {
		code, _ = strconv.Atoi(ret)
		bool2 = true
	}
	return code, bool2
}
func SendMessage(tel string, code string) error {
	ret := ConnectRedis.Get(context.TODO(), fmt.Sprintf("Code:telc:%s", tel))
	if ret == nil {
		return fmt.Errorf("验证码已发送,稍后在试")
	}
	//参数一：连接的节点地址（有很多节点选择，这里我选择杭州）
	//参数二：AccessKey ID
	//参数三：AccessKey Secret
	//client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "AccessKey ID", "AccessKey Secret")
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI5tRZK93Zqe9rVv7YitF7", "6GxZ2zxKcUKO63MIjv61lHlEffUDzq")
	request := dysmsapi.CreateSendSmsRequest()       //创建请求
	request.Scheme = "https"                         //请求协议
	request.PhoneNumbers = tel                       //接收短信的手机号码
	request.SignName = "阿里云短信测试"                     //短信签名名称
	request.TemplateCode = "SMS_154950909"           //短信模板ID
	par, err := json.Marshal(map[string]interface{}{ //定义短信模板参数（具体需要几个参数根据自己短信模板格式）
		"code": code,
	})
	request.TemplateParam = string(par)      //将短信模板参数传入短信模板
	response, err := client.SendSms(request) //调用阿里云API发送信息
	if err != nil {                          //处理错误
		fmt.Print(err.Error())
	}
	_ = ConnectRedis.Set(context.TODO(), fmt.Sprintf("Code:tel:%s", tel), code, 600*time.Second)
	_ = ConnectRedis.Set(context.TODO(), fmt.Sprintf("Code:telc:%s", tel), 100, 60*time.Second)
	fmt.Printf("response is %#v\n", response) //控制台输出响应
	return nil
}
func GetMessage(tel string) (code string) {
	code, _ = ConnectRedis.Get(context.TODO(), fmt.Sprintf("Code:tel:%s", tel)).Result()
	return code
}
