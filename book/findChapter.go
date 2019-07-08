package book

import (
	"fmt"
	"github.com/novel/sq"
	"github.com/novel/bookType"
	"io/ioutil"
)

//查找书籍所有章节，返回章节对象list，参数（书籍ID）
func FindAllChapter(bookNumber int) []bookType.Chapter{
	var chapters []bookType.Chapter
	db := sq.SqConnection()
	rows, err := db.Query(" select * from chapter where book_id = ? ;", bookNumber )
	if err != nil {
		fmt.Println("err: ", err)
	}

	for rows.Next() {
		var chapter bookType.Chapter = bookType.Chapter{}
		rows.Scan(&chapter.ChaId, &chapter.ChaName, &chapter.ChaSize, &chapter.BookId, &chapter.Path)
		chapters = append(chapters, chapter)
	}
	return chapters
}

//查找书籍具体章节内容，返回章节内容，参数（书籍ID，章节ID）
func FindChapter(bookId, chapterNumber int) string{
	var a bookType.Chapter = bookType.Chapter{}
	db := sq.SqConnection()
	rows, err := db.Query(" select * from chapter where book_id = ? and chapter_id = ?;", bookId, chapterNumber)
	if err != nil {
		fmt.Println("err: ", err)
	}
	rows.Next()
	rows.Scan(&a.ChaId, &a.ChaName, &a.ChaSize, &a.BookId, &a.Path)
	chapterData := ReadChaper(a.Path)
	IncreaseReadingNumber(bookId)
	return chapterData

}

//读取章节内容， 返回章节内容， 参数（章节路径地址）
func ReadChaper(path string) string {
	var chapterData, err= ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(chapterData)
}