package controllers

import (
	"github.com/astaxie/beego"
)

type ChangePasswordController struct {
	beego.Controller
}

type SignUpController struct {
	beego.Controller
}

type ChangeUsernameByTokenController struct {
	beego.Controller
}

func (c *ChangePasswordController) Put() {
	username := c.GetString("username")
	oldPassword := c.GetString("oldpassword")
	newPassword := c.GetString("newpassword")
	emailAdd := c.GetString("email")
	emailCode, _ := c.GetInt("emailcode")
	if changePassword(username, oldPassword, newPassword, emailAdd, emailCode) {
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

func (c *ChangeUsernameByTokenController) Put() {
	token := c.GetString("token")
	newUsername := c.GetString("newusername")
	if changeUsernameByToken(token, newUsername) {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':401,'massage':'failed'}")
	}
}
