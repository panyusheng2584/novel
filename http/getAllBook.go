package http

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/novel/book"
)

type AllBookController struct {
	beego.Controller
}


func (this *AllBookController) Get() {

	bookList := book.FindAllBook()

	bookStr, _ := json.Marshal(bookList)

	this.Ctx.WriteString(string(bookStr))
}