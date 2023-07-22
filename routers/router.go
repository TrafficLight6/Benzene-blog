package routers

import (
	"benzenz-blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/add/page", &controllers.AddPageController{})
	beego.Router("/add/signup", &controllers.SignUpController{})
	beego.Router("/add/token/createtoken", &controllers.CreateTokenController{})
	beego.Router("/add/token/addpage", &controllers.AddPageByTokenController{})
	beego.Router("/del/page", &controllers.DelPageController{})
	beego.Router("/del/token/page", &controllers.DelPageByTokenController{})
	beego.Router("/del/token/deltoken", &controllers.DelTokenController{})
	beego.Router("/edit/page", &controllers.EditPageController{})
	beego.Router("/edit/token/page", &controllers.EditPageByTokenController{})
	beego.Router("/get/admincheck", &controllers.AdminCheckController{})
	beego.Router("/get/username", &controllers.GetUsernameController{})
	beego.Router("/get/pagejson", &controllers.GetPageListController{})
	beego.Router("/get/page", &controllers.GetPageMainController{})
	beego.Router("/edit/password", &controllers.ChangePasswordController{})

	beego.Router("email/send", &controllers.SendEmailController{})
}
