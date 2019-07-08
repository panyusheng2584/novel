package http

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/novel/book"
	"strconv"
)

type AllChapetrController struct {
	beego.Controller
}

type bookId struct {
	Code string `json:"code"`
}



func (this *AllChapetrController)Post() {

	var ob bookId
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil{
		fmt.Println("err == ", err)
	}

	int, err := strconv.Atoi(ob.Code)
	chapterList := book.FindAllChapter(int)
	chapterJson, _ := json.Marshal(chapterList)
	this.Ctx.WriteString(string(chapterJson))
}
