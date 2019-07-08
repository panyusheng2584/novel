package http

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/novel/user"
)

type UserInfoController struct {
	beego.Controller
}


func (this *UserInfoController)Get() {
	userName := this.Ctx.GetCookie("name")
	user := user.UserInfo(userName)

	classificationJson, _ := json.Marshal(user)

	this.Ctx.WriteString(string(classificationJson))
}