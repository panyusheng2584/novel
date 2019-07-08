package user

import (
	"fmt"
	"github.com/novel/sq"
	"github.com/novel/bookType"
)

//用户登陆验证，返回0，1，2，0代表用户不存在，1代表用户验证成功，2代表用户密码错误， 参数（用户名，用户密码）
func UserLogin (name, password string) int {
	var user bookType.SqUser
	db := sq.SqConnection()
	rows, err := db.Query("SELECT * FROM user WHERE user_name = ? ;",(name))
	if err != nil{
		fmt.Println("err :", err)
	}
	if rows.Next(){
		rows.Scan(&user.Name, &user.Password, &user.LikeBook, &user.ReadBook, &user.Money)
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

