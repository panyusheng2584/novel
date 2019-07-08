package book

import (
	"fmt"
	"github.com/novel/sq"
)

//查找书籍分类， 返回书籍分类list
func Classification() []string{
	var books []string
	var a string
	db := sq.SqConnection()
	rows, err := db.Query("select  distinct  book_classification  from  books;")
	if err != nil {
		fmt.Println("err: ", err)
	}

	for rows.Next() {
		rows.Scan(&a)
		books = append(books, a)
	}
	return books

}