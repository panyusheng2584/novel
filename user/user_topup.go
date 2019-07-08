package user

import (
	"fmt"
	"github.com/novel/sq"
)

func UserTopup (UserName string, password int)  bool{
	if password != 111111 {
		fmt.Println("jia")
		return false
	}
	db := sq.SqConnection()
	num := 100
	_, err := db.Query("UPDATE user SET money = money+? where user_name=?;", num, UserName)
	if err != nil {
		fmt.Println("err :", err)
		return false
	}
	return true
}


func UserDeductMoney (UserName string, num int) bool{
	db := sq.SqConnection()
	_, err := db.Query("UPDATE user SET money = money-? where user_name=?;", num, UserName)
	if err != nil {
		fmt.Println("err :", err)
		return false
	}
	return true

}