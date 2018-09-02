package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct{
	beego.Controller
}

func (c *UserController) Get() {
	beego.Info("user log test...................")
	c.Data["user"] = "edison"
	c.TplName = "user.html"
}
