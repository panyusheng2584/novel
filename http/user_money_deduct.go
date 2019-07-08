package http

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/novel/user"
)

type UserMoneyDeductController struct {
	beego.Controller
}

type DeductNum struct {
	Num int `json:"num"`
}

func (this *UserMoneyDeductController)Post() {
	userName := this.Ctx.GetCookie("name")
	var ob DeductNum
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil{
		fmt.Println("err == ", err)
	}

	judeMoney := user.UserDeductMoney(userName, ob.Num)
	fmt.Println(judeMoney)

	classificationJson, _ := json.Marshal(judeMoney)

	this.Ctx.WriteString(string(classificationJson))
}