package http

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/novel/book"
)

type ClassificationListController struct {
	beego.Controller
}


func (this *ClassificationListController)Get() {

	classificationList := book.Classification()

	classificationJson, _ := json.Marshal(classificationList)

	this.Ctx.WriteString(string(classificationJson))
}