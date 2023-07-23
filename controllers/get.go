package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

type AdminCheckController struct {
	beego.Controller
}

type GetUsernameController struct {
	beego.Controller
}

type GetPageListController struct {
	beego.Controller
}

type GetPageMainController struct {
	beego.Controller
}

type GetUserIdByTokenController struct {
	beego.Controller
}

type Pagelist struct {
	Idlist     []int    `json:"id"`
	Namelist   []string `json:"name"`
	Posterlist []string `json:"poster"`
	Timelist   []int    `json:"time"`
}

func (c *AdminCheckController) Get() {
	username := c.GetString("username")
	password := c.GetString("password")
	if adminCheck(username, password) {
		c.Ctx.WriteString("{'code':200,'massage':'admin'}")
	} else {
		c.Ctx.WriteString("{'code':200,'massage':'user or not found'}")
	}
}

func (c *GetUsernameController) Get() {
	var output string
	id, err := c.GetInt("id")
	if err != nil {
		output = "{'code':404,'massage':'failed'}"
	} else {
		output = "{'code':200,'massage':'" + getUsername(id) + "'}"
	}
	c.Ctx.WriteString(output)
}

func (c *GetPageListController) Get() {
	var pagelist Pagelist
	least, erra := c.GetInt("least")
	most, errb := c.GetInt("most")
	var outputJson string
	if erra == nil && errb == nil {
		pagelist.Idlist, pagelist.Namelist, pagelist.Posterlist, pagelist.Timelist = getPageList(least, most)
		jsonByte, _ := json.Marshal(pagelist)
		listJson := string(jsonByte)
		outputJson = "{'code':200,'massage':" + listJson + "}"
	} else {
		outputJson = "{'code':404,'massage':'failed'}"
	}
	c.Ctx.WriteString(outputJson)
}

func (c *GetPageMainController) Get() {
	var output string
	id, err := c.GetInt("page_id")
	if err != nil {
		output = "{'code':'404','massage':'none'}"
	} else {
		text := getPage(id)
		if text == "none" {
			output = "{'code':'404','massage':'none'}"
		} else {
			output = "{'code':'200','massage':'" + text + "'}"
		}
	}
	c.Ctx.WriteString(output)
}

func (c *GetUserIdByTokenController) Get() {
	id, _ := getUserIdByToken(c.GetString("token"))
	if id != -1 {
		c.Ctx.WriteString("{'code':'200','massage':" + strconv.Itoa(id) + "}")
	} else {
		c.Ctx.WriteString("{'code':404,'massage':'failed'}")
	}
}
