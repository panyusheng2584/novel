package user

import (
	"fmt"
	"github.com/novel/bookType"
	"github.com/novel/sq"
)
type userInfo struct {
	BookName string	`json:"bookName"`
	BookId	string `json:"bookId"`
	ChapterId string`json:"chapterId"`
}


//获取用户信息， 返回用户对象，参数（用户名）
func UserInfo (userName string) bookType.User {
	var sqUser bookType.SqUser
	var user bookType.User
	var read userInfo
	var like userInfo
	db := sq.SqConnection()
	rows, err := db.Query("SELECT user_name, user_password, moeny FROM user WHERE user_name = ? ;",(userName))
	if err != nil{
		fmt.Println("err :", err)
	}
	rows.Next()
	rows.Scan(&sqUser.Name, &sqUser.Password)


	rows, err = db.Query("SELECT bookID, bookName FROM userRead WHERE userName = ? ;",(userName))
	if err != nil{
		fmt.Println("err :", err)
	}
	for rows.Next(){
		rows.Scan(&read.BookId, &read.BookName)
		var readBook = []string{read.BookName, read.BookId}
		user.Read = append(user.Read, readBook)
	}

	rows, err = db.Query("SELECT bookID, bookName FROM userLike WHERE userName = ? ;",(userName))
	if err != nil{
		fmt.Println("err :", err)
	}
	for rows.Next(){
		rows.Scan(&like.BookId, &like.BookName)
		var likeBook = []string{read.BookName, read.BookId}
		user.Read = append(user.Like, likeBook)
	}

	return user
}

func UserMoney (userName string) int {
	var money int
	db := sq.SqConnection()
	rows, err := db.Query("SELECT money FROM user WHERE user_name = ? ;",(userName))
	if err != nil{
		fmt.Println("err :", err)
	}
	rows.Next()
	rows.Scan(&money)
	return money

}