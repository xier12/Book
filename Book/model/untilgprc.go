package model

import (
	"Book/pb/login"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
	"time"
)

func UntilGPRCClient() {

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
		Name:         "root",
		Password:     "root",
		CaptchaId:    "",
		CaptchaValue: "",
	})
	fmt.Println(result, err)
}
func EndVote() {
	orders := make([]Order, 0)

	if err := ConnectMysql.Raw("select * from order where mybool = ?", 0).Scan(&orders).Error; err != nil {
		return
	}
	now := time.Now().Unix()

	for _, order := range orders {
		num, _ := strconv.ParseInt(order.Createtime, 10, 64)
		if num+600 <= now {
			ConnectMysql.Exec("update order set during = ? where id = ? ", 1, order.ID)
		}
	}
	return
}
