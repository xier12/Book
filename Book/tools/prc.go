package tools

import (
	"Book/pb/login"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func PRC(username, userpassword, usrcaptid, usercaptvalue string) *login.Response {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	defer func() {
		_ = conn.Close()
	}()

	// 初始化客户端
	client := login.NewLoginByUserNameServiceClient(conn)

	result, err := client.LoginByUserName(context.Background(), &login.Request{
		Name:         username,
		Password:     userpassword,
		CaptchaId:    usrcaptid,
		CaptchaValue: usercaptvalue,
	})
	fmt.Println(result, err)
	return result
}
