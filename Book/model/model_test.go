package model

import (
	"fmt"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {
	CoreConf()
	ConnectMysql = UntilMysql()
	fmt.Println(333)
	var u User
	u = User{
		Id:       1,
		Uuid:     0,
		Name:     "aaaaaaa",
		Password: "abcdrfg",
		Ctime:    time.Now().Format("2006-01-02 15:04:05"),
		UTime:    time.Now().Format("2006-01-02 15:04:05"),
		Power:    "0",
	}
	AddUser(u)
	fmt.Println(22222)
}
