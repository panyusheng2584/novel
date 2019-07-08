package http

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/novel/user"
)

type UserValidationController struct {
	beego.Controller
}


type userdata struct {
	Name string	`json:"name"`
	Password string	`json:"password"`
}


func (this *UserValidationController) Post() {
	var ob userdata
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil{
		fmt.Println("err == ", err)
	}
	a := user.UserLogin(ob.Name, ob.Password)
	if a == 1{
		this.Ctx.SetCookie("name", ob.Name, 1000, "/")  // 设置cookie
		this.Ctx.SetCookie("password", ob.Password, 1000, "/")  // 设置cookie
	}
	judgeUser, _ := json.Marshal(a)
	this.Ctx.WriteString(string(judgeUser))

}


/*
func (c *UserValidationController) Login() {
	name := c.Ctx.GetCookie("name")
	password := c.Ctx.GetCookie("password")

	if name != "" && password != ""{
		c.Ctx.Redirect(302, "/books.html")
	} else {
		formData := `<html><form action="/login" method="post">
         <input type="text" name="Name">
         <input type="password" name="Password">
        <input type="submit" value="post">
            </html>
        `
		c.Ctx.WriteString(formData)
	}
}



func (c *UserValidationController) PostData() {

	user := bookType.Userdata{}
	if err := c.ParseForm(&user); err != nil {

	}
	a := books.UserLogin(user.Name, user.Password)
	if a == 1{
		c.Ctx.SetCookie("name", user.Name, 1000, "/")  // 设置cookie
		c.Ctx.SetCookie("password", user.Password, 1000, "/")  // 设置cookie
		c.Ctx.Redirect(302, "/books.html")
	}
	if a == 2{
		formData := `<html><p>密码错误</p></br><form action="/login" method="post">
         <input type="text" name="Name">
         <input type="password" name="Password">
        <input type="submit" value="post">
            </html>
        `
		c.Ctx.WriteString(formData)
	}

	if a == 0{
		formData := `<html><p>用户名不存在</p></br><form action="/login" method="post">
         <input type="text" name="Name">
         <input type="password" name="Password">
        <input type="submit" value="post">
		<a href="registration.html">注册</a>
            </html>
        `
		c.Ctx.WriteString(formData)
	}

}
*/
