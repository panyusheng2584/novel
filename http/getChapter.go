package http

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/novel/book"
	"github.com/novel/user"
	"strconv"
)

type ChapetrController struct {
	beego.Controller
}

type chapterInfo struct {
	BookName string	`json:"bookName"`
	BookId	string `json:"bookId"`
	ChapterId string`json:"chapterId"`
}

type UserInfo struct {
	Chapter string `json:"chapter"`
	UserMoney int `json:"userMoney"`
}

func (this *ChapetrController)Post() {
	name := this.Ctx.GetCookie("name")
	var userData UserInfo
	var ob chapterInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil{
		fmt.Println("err == ", err)
	}
	fmt.Println(ob)
	bookId, err := strconv.Atoi(ob.BookId)
	chaId, err := strconv.Atoi(ob.ChapterId)

	userData.Chapter = book.FindChapter(bookId, chaId)
	userData.UserMoney = user.UserMoney(name)
	user.UserReadBookIncrease(name, ob.BookName, bookId)

	chapterJson, _ := json.Marshal(userData)
	this.Ctx.WriteString(string(chapterJson))
}

