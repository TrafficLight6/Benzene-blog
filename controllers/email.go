package controllers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	mail "github.com/xhit/go-simple-mail"
)

type EmailCode struct {
	Id    int    `db:"id"`
	Email string `db:"email"`
	Code  int    `db:"code"`
}

type SendEmailController struct {
	beego.Controller
}

type EmailCheckController struct {
	beego.Controller
}

func sendEmail(emailadd string) bool {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed   :", err, "\n")
		return false
	}
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(89999) + 10000
	sql := "INSERT INTO blog_email_code(id,email,code) VALUE (DEFAULT,?,?)"
	r, err := db.Exec(sql, emailadd, code)
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		return false
	}
	_, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		conn.Rollback()
		return false
	}
	server := mail.NewSMTPClient()
	server.Host = "smtp.126.com"
	server.Port = 25
	server.Username = "TrafficLight6@126.com"
	server.Password = "SUNFRNSQZFUAXKCK"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		fmt.Println(err)
		conn.Rollback()
		return false
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("From Me <TrafficLight6@126.com>")
	email.AddTo(emailadd)
	email.SetSubject("Email Code")

	htmlStr := "<h1>your code is " + strconv.Itoa(code) + "</h1>"
	email.SetBody(mail.TextHTML, htmlStr) //发送html信息
	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		fmt.Println(err)
		conn.Rollback()
		return false
	}
	conn.Commit()
	return true
}

func checkEmail(emailadd string, emailcode int) bool {
	var ec []EmailCode
	sql := "SELECT id,email,code FROM blog_email_code WHERE email=? AND code=? "
	err := db.Select(&ec, sql, emailadd, emailcode)
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		return false
	}
	if ec == nil {
		return false
	}
	return true
}

func (c *SendEmailController) Post() {
	if sendEmail(c.GetString("email")) {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':404,'massage':'faileds'}")
	}
}

func (c *EmailCheckController) Get() {
	code, err := c.GetInt("code")
	if (checkEmail(c.GetString("email"), code)) && (err == nil) {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':404,'massage':'faileds'}")
	}
}
