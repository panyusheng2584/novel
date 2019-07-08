package main

import (
	"github.com/astaxie/beego"
	"github.com/novel/http"
)

func main() {
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.SetStaticPath("/", "web")
	beego.Router("/books", &http.AllBookController{})
	beego.Router("/book", &http.BookController{})
	beego.Router("/chapters", &http.AllChapetrController{})
	beego.Router("/chapter", &http.ChapetrController{})
	beego.Router("/classification", &http.ClassificationListController{})
	beego.Router("/login", &http.UserValidationController{})
	//beego.Router("/login", &http.UserValidationController{}, "post:PostData;get:Login")
	beego.Router("/registration", &http.UserRegistrationController{})
	beego.Router("/getUser", &http.UserInfoController{})
	beego.Router("/adminlogin", &http.AdminValidationController{})
	beego.Router("/adminregistration", &http.AdminRegistrationController{})
	beego.Router("/usertopup", &http.UserTopupController{})
	beego.Router("/userdeduct", &http.UserMoneyDeductController{})
	beego.Router("/classbooks", &http.AllClassBookController{})
	beego.Run()
}

