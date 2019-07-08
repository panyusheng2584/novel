package http

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/novel/book"
)

type BookController struct {
	beego.Controller
}

type bookName struct {
	Name string		`json:"name"`
}

func (this *BookController)Post() {
	name := bookName{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &name)
	if err != nil{
		fmt.Println("err :", err)
	}

	book := book.FindBook(name.Name)
	bookStr, _ := json.Marshal(book)

	this.Ctx.WriteString(string(bookStr))
}