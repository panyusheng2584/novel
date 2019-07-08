package http

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/novel/bookType"
	"github.com/novel/user"
)

type UserRegistrationController struct {
	beego.Controller
}



func (this *UserRegistrationController)Post() {
	var ob bookType.Userdata
	if err := this.ParseForm(&ob); err != nil {
		fmt.Println("err :", err)
	}
	this.Ctx.SetCookie("name", ob.Name, 1000, "/")  // 设置cookie
	this.Ctx.SetCookie("password", ob.Password, 1000, "/")
	data := user.UserRegistered(ob.Name, ob.Password)

	if data == 1{
		this.Ctx.Redirect(302, "/books.html")
	}
}
