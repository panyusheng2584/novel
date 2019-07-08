package http

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/novel/admin"
)

type AdminValidationController struct {
	beego.Controller
}


type admindata struct {
	Name string	`json:"name"`
	Password string	`json:"password"`
}


func (this *AdminValidationController) Post() {
	fmt.Println("diaoyongl")
	var ob admindata
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil{
		fmt.Println("err == ", err)
	}
	a := admin.AdminLogin(ob.Name, ob.Password)
	if a == 1{
		this.Ctx.SetCookie("name", ob.Name, 1000, "/")  // 设置cookie
		this.Ctx.SetCookie("password", ob.Password, 1000, "/")  // 设置cookie
		this.Ctx.SetCookie("class","admin",1000, "/")
	}
	fmt.Println(a)
	judgeUser, _ := json.Marshal(a)
	this.Ctx.WriteString(string(judgeUser))

}
