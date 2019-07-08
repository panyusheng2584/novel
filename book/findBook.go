package book

import (
	"fmt"
	"github.com/novel/bookType"
	"github.com/novel/sq"
)

//查找所有书籍， 返回书籍对象list
func FindAllBook() []bookType.Book {
	var books []bookType.Book
	db := sq.SqConnection()
	rows, err := db.Query("select * from books;")
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

//查找所有书籍， 返回书籍对象 ， 参数（书籍名字）
func FindBook(bookName string) bookType.Book {

	db := sq.SqConnection()
	rows, err := db.Query(" select * from books where book_name = ?;", (bookName))
	if err != nil {
		fmt.Println("err: ", err)
	}

	rows.Next()
	var book bookType.Book = bookType.Book{}
	rows.Scan(&book.BookId, &book.BookName, &book.BookSize, &book.LikeNumber, &book.ReadingNumber, &book.BookPath, &book.ChapterNumber, &book.BookAuthor, &book.BookClassification, &book.Description)
	return book
}