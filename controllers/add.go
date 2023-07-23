package controllers

import (
	"github.com/astaxie/beego"
)

type AddPageController struct {
	beego.Controller
}

type AddPageByTokenController struct {
	beego.Controller
}

func (c *AddPageController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	title := c.GetString("title")
	main := c.GetString("main")
	if addPage(username, password, title, main) {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':401,'massage':'failed'}")
	}
}

func (c *AddPageByTokenController) Post() {
	token := c.GetString("token")
	title := c.GetString("title")
	main := c.GetString("main")
	if addPageByToken(token, title, main) {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':401,'massage':'failed'}")
	}
}
