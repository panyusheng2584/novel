package bookType


type Book struct {
	BookId int	`json:"bookId"`
	BookName string	`json:"bookName"`
	BookSize int	`json:"bookSize"`
	LikeNumber int	`json:"likeNumber"`
	ReadingNumber int	`json:"readingNumber"`
	BookPath string	`json:"bookPath"`
	ChapterNumber int	`json:"ChapterNumber"`
	BookAuthor string	`json:"bookAuthor"`
	BookClassification string	`json:"bookClassification"`
	Description string	`json:"description"`
}