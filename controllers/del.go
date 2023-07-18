package controllers

import (
	"github.com/astaxie/beego"
)

type DelPageController struct {
	beego.Controller
}

type DelPageByTokenController struct {
	beego.Controller
}

type DelTokenController struct {
	beego.Controller
}

func (c *DelPageController) Delete() {
	username := c.GetString("username")
	password := c.GetString("password")
	id, err := c.GetInt("page_id")
	if delPage(username, password, id) && err == nil {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':401,'massage':'failed'}")
	}
}

func (c *DelPageByTokenController) Delete() {
	token := c.GetString("token")
	id, err := c.GetInt("page_id")
	if delPageByToken(token, id) && err == nil {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':401,'massage':'failed'}")
	}
}

func (c *DelTokenController) Delete() {
	token := c.GetString("token")
	if delToken(token) {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':401,'massage':'failed'}")
	}
}
