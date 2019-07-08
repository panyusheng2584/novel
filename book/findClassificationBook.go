package book

import (
	"fmt"
	"github.com/novel/bookType"
	"github.com/novel/sq"
)

//分类查找书籍，返回书籍对象list，参数（书籍分类名字）
func FindClassification(classificationName string) []bookType.Book{
	var books []bookType.Book
	db := sq.SqConnection()
	rows, err := db.Query("select * from books where book_classification = ?;", (classificationName))
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