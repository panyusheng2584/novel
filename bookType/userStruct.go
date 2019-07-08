package bookType


type SqUser struct {
	Name string 	`json:"name"`
	Password string	`json:"password"`
	LikeBook string	`json:"likeBook"`
	ReadBook string	`json:"readBook"`
	Money int		`json:"money"`
}

type User struct {
	Name string
	Like [][]string
	Read [][]string
	BookId string
}


type Userdata struct {
	Name string		`json:"name"`
	Password string	`json:"password"`
}