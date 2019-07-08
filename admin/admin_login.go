package admin

import (
	"fmt"
	"github.com/novel/bookType"
	"github.com/novel/sq"
)

func AdminLogin (name, password string) int {
	var user bookType.SqUser
	db := sq.SqConnection()
	rows, err := db.Query("SELECT * FROM admin WHERE admin_name = ? ;",(name))
	if err != nil{
		fmt.Println("err :", err)
	}
	if rows.Next(){
		rows.Scan(&user.Name, &user.Password)
		if user.Password == password{
			return 1
		}else {
			fmt.Println("password err")
			return 2
		}
	}else {
		fmt.Println("user err")
		return 0
	}

}

