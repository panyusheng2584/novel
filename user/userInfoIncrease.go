package user

import (
	"fmt"
	"github.com/novel/book"
	"github.com/novel/sq"
)


//增加用户点赞记录， 参数（用户名， 书籍ID）
func UserLikeBookIncrease(UserName string, bookId int)  {
	db := sq.SqConnection()
	rows, err := db.Query("SELECT * FROM userLike WHERE userName = ? and bookId = ?;", UserName, bookId)
	if err != nil {
		fmt.Println("err :", err)
	}
	if rows.Next() {
		return
	}
	rows, err = db.Query("INSERT INTO userLike (userName, bookId) VALUES (?, ?);", UserName, bookId)
	book.IncreaseLikeNumber(bookId)

}

//增加用户阅读记录, 参数（用户名， 书籍ID）
func UserReadBookIncrease(UserName, bookName string, bookId int)  {
	db := sq.SqConnection()
	rows, err := db.Query("SELECT * FROM userRead WHERE userName = ? and bookId = ?;", UserName, bookId)
	if err != nil{
		fmt.Println("errr  :", err)
	}
	if rows.Next() {
		return
	}
	rows, err = db.Query("INSERT INTO userRead (userName, bookId, bookName) VALUES (?, ?, ?);", UserName, bookId, bookName)
	if err != nil{
		fmt.Println("errr  :", err)
	}
}
