package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// 最基础的版本
func Encrypt(pwd string) string {

	hash := md5.New()
	hash.Write([]byte(pwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	//fmt.Printf("加密后的密码：%s\n", hashString)

	return hashString
}

func EncryptV1(pwd string) string {
	//最常用,用户密码+自己设计的密钥
	newPwd := pwd + "Lexington" //不能随便起，且不能暴露
	hash := md5.New()
	hash.Write([]byte(newPwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	//fmt.Printf("加密后的密码：%s\n", hashString)

	return hashString
}

// 最安全的版本
func EncryptV2(pwd string) string {
	//基于Blowfish 实现加密。简单快速，但有安全风险
	//golang.org/x/crypto/ 中有大量的加密算法
	newPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("密码加密失败：", err)
		return ""
	}
	newPwdStr := string(newPwd)
	fmt.Printf("加密后的密码：%s\n", newPwdStr)
	return newPwdStr
}
