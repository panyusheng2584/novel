package book

import (
	"fmt"
	"github.com/novel/sq"
	"github.com/novel/bookType"
)

//查找书籍字数排行，返回书籍对象list
func NumberList() []bookType.Book{
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
