package controllers

import (
	"github.com/astaxie/beego"
)

type CreateTokenController struct {
	beego.Controller
}

type DelTokenController struct {
	beego.Controller
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

func (c *DelTokenController) Delete() {
	token := c.GetString("token")
	if delToken(token) {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':401,'massage':'failed'}")
	}
}
