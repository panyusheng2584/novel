package bookType


type Chapter struct {
	ChaId   int		`json:"chaId"`
	ChaName string	`json:"chaName"`
	ChaSize int		`json:"chaSize"`
	BookId int		`json:"bookId"`
	Path string		`json:"path"`
}