package book

import (
	"fmt"
	"github.com/novel/sq"
	"github.com/novel/bookType"
)

//查找书籍阅读排行，返回书籍list
func ReadingList() []bookType.Book{
	var books []bookType.Book
	db := sq.SqConnection()
	rows, err := db.Query("select * from books order by book_size DESC ;")
	if err != nil {
		fmt.Println("err: ", err)
	}

	for rows.Next() {
		var book bookType.Book = bookType.Book{}
		rows.Scan(&book.BookId, &book.BookName, &book.BookSize, &book.LikeNumber, &book.ReadingNumber, &book.BookPath, &book.ChapterNumber, &book.BookAuthor, &book.BookClassification, &book.Description)
		books = append(books, book)
	}

	return books

}

//增加书籍阅读次数+1，参数（书籍ID）
func IncreaseReadingNumber(book_id int) {
	db := sq.SqConnection()
	_, err := db.Query("UPDATE books SET reading_number = reading_number+1 WHERE book_id = ? ;", (book_id))
	if err != nil{
		fmt.Println("err :",err)
	}

}