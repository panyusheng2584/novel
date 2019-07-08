package admin

import (
	"fmt"
	"github.com/novel/sq"
)
func AdminRegistered (name, password string) int {
	db := sq.SqConnection()
	_, err := db.Query("INSERT INTO admin (admin_name, password )  VALUES  (?, ? ) ;",name, password)
	if err != nil{
		fmt.Println("err :", err)
		return 0
	}else {
		fmt.Println("Registered successfully")
		return 1
	}

	return 0

}