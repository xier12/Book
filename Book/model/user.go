package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func FindUserByName(uname string) User {
	var user User
	//a := GetRedis("User:name:" + uname)
	//err := json.Unmarshal([]byte(a), &user)
	//if err != nil {
	//	fmt.Println(err)
	//}
	if user.Id > 0 {
		return user
	}
	ConnectMysql.Raw("select * from user where name = ?", uname).Scan(&user)
	JsonUser, err1 := json.Marshal(user)
	err2 := SetRedis("User:name:"+user.Name, JsonUser, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return user
}
func FindUserById(uid int64) User {
	var user User
	a := GetRedis("User:id:" + strconv.FormatInt(uid, 10))
	err := json.Unmarshal([]byte(a), &user)
	if err != nil {
		fmt.Println(err)
	}
	ConnectMysql.Raw("select * from user where id = ?", uid).Scan(&user)
	JsonUser, err1 := json.Marshal(user)
	err2 := SetRedis("User:id:"+strconv.FormatInt(uid, 10), JsonUser, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return user
}
func FindUserByTel(utel string) User {
	var user User
	//a := GetRedis("User:tel:" + utel)
	//err := json.Unmarshal([]byte(a), &user)
	//if err != nil {
	//	fmt.Println(err)
	//}
	ConnectMysql.Raw("select * from user where tel = ?", utel).Scan(&user)
	JsonUser, err1 := json.Marshal(user)
	err2 := SetRedis("User:tel:"+utel, JsonUser, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return user
}
func FindUserByEmail(uemail string) User {
	var user User
	//a := GetRedis("User:email:" + uemail)
	//err := json.Unmarshal([]byte(a), &user)
	//if err != nil {
	//	fmt.Println(err)
	//}
	ConnectMysql.Raw("select * from user where email = ?", uemail).Scan(&user)
	JsonUser, err1 := json.Marshal(user)
	err2 := SetRedis("User:email:"+uemail, JsonUser, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return user
}
func AddUser(user User) bool {
	err := ConnectMysql.Table("user").Create(&user).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	JsonUser, err1 := json.Marshal(user)
	err2 := SetRedis("User:name:"+user.Name, JsonUser, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return true
}
