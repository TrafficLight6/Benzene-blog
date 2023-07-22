package controllers

import (
	"github.com/astaxie/beego"
)

type AddPageController struct {
	beego.Controller
}

type CreateTokenController struct {
	beego.Controller
}

type AddPageByTokenController struct {
	beego.Controller
}

type SignUpController struct {
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

func (c *CreateTokenController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	token := createToken(username, password)
	var returnjson string
	if token == "none" {
		returnjson = "{'code':404,'massage':'error'}"
	} else {
		returnjson = "{'code':200,'massage':'" + token + "'}"
	}
	c.Ctx.WriteString(returnjson)
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

func (c *SignUpController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	email := c.GetString("email")
	code, _ := c.GetInt("emailcode")
	if signUp(username, password, email, code) {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':401,'massage':'failed'}")
	}
}
