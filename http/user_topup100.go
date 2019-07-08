package http

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/novel/user"
)

type UserTopupController struct {
	beego.Controller
}

type TopupPassword struct {
	Code int `json:"code"`
}

func (this *UserTopupController)Post() {
	userName := this.Ctx.GetCookie("name")
	var ob TopupPassword
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil{
		fmt.Println("err == ", err)
	}

	judeMoney := user.UserTopup(userName, ob.Code)
	fmt.Println(ob)

	classificationJson, _ := json.Marshal(judeMoney)

	this.Ctx.WriteString(string(classificationJson))
}