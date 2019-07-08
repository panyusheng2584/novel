package http

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/novel/book"
)

type AllClassBookController struct {
	beego.Controller
}

type ClassBookName struct {
	Name string `json:"name"`
}



func (this *AllClassBookController)Post() {

	var ob ClassBookName
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil{
		fmt.Println("err == ", err)
	}
	fmt.Println(ob)

	classBookList := book.FindClassification(ob.Name)
	fmt.Println(classBookList)
	classBookListJson, _ := json.Marshal(classBookList)
	this.Ctx.WriteString(string(classBookListJson))
}