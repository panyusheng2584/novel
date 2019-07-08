package user

import (
	"fmt"
	"github.com/novel/sq"
)

//用户注册， 返回0，1，0注册失败，1注册成功， 参数（用户名，用户密码）
func UserRegistered (name, password string) int {
	db := sq.SqConnection()
	_, err := db.Query("INSERT INTO user (user_name, user_password, like_book, read_book, money )  VALUES  (?, ?, 0, 0, 0 ) ;",name, password)
	if err != nil{
		fmt.Println("err :", err)
		return 0
	}else {
		fmt.Println("Registered successfully")
		return 1
	}

	return 0

}
