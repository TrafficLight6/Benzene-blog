package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type EditPageController struct {
	beego.Controller
}

type EditPageByTokenController struct {
	beego.Controller
}

func (c *EditPageController) Put() {
	username := c.GetString("username")
	password := c.GetString("password")
	id, err := c.GetInt("page_id")
	main := c.GetString("main")
	if editPage(username, password, id, main) && err == nil {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':401,'massage':'failed'}")
	}
}

func (c *EditPageByTokenController) Put() {
	token := c.GetString("token")
	id, err := c.GetInt("page_id")
	main := c.GetString("main")
	if editPageByToken(token, id, main) && err == nil {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		fmt.Println(err, "\n")
		c.Ctx.WriteString("{'code':401,'massage':'failed'}")
	}
}
