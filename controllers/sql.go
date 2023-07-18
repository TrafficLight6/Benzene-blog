package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type user struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Pwd      string `db:"password"`
	Add      string `db:"allowadd"`
	Post     string `db:"allowpost"`
	Admin    string `db:"admin"`
	Email    string `db:"email"`
}

type page struct {
	Id     int    `db:"id"`
	UserId int    `db:"user_id"`
	Title  string `db:"title"`
	Main   string `db:"main"`
	Time   int    `db:"time"`
}

type tokenType struct {
	Id        int    `db:"id"`
	UserId    int    `db:"user_id"`
	MainToken string `db:"main_token"`
	Time      int    `db:"time"`
}

var db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "bloguser:blogpwd@tcp(127.0.0.1:3306)/blogdb")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}

func signUp(username string, password string, email string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	bytes := hash.Sum(nil)
	sum := hex.EncodeToString(bytes)

	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return false
	}
	sql := "INSERT INTO blog_user(id, username,password,allowadd,allowpost,admin,email) VALUES(DEFAULT,?,?,DEFAULT,DEFAULT,DEFAULT,?)"
	r, err := db.Exec(sql, username, sum, email)
	if err != nil {
		fmt.Println("exec failed   :", err)
		conn.Rollback()
		return false
	}
	_, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed   :", err)
		conn.Rollback()
		return false
	}
	conn.Commit()
	return true

}

func addAllow(username string, password string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	bytes := hash.Sum(nil)
	sum := hex.EncodeToString(bytes)
	var users []user
	sql := "SELECT id, username,password,allowadd,allowpost,admin,email FROM blog_user WHERE username=? AND password=? "
	err := db.Select(&users, sql, username, sum)
	if err != nil {
		fmt.Printf("user select failed   :", err, "\n")
		return false
	}
	if users == nil {
		return false
	}
	for index, _ := range users {
		if users[index].Add == "false" {
			return false
		}
	}
	return true
}

func postAllow(username string, password string) bool { //评论功能未开发
	hash := sha256.New()
	hash.Write([]byte(password))
	bytes := hash.Sum(nil)
	sum := hex.EncodeToString(bytes)
	var users []user
	sql := "SELECT id, username,password,allowadd,allowpost,admin,email FROM blog_user WHERE username=? AND password=? "
	err := db.Select(&users, sql, username, sum)
	if err != nil {
		fmt.Printf("user select failed   :", err, "\n")
		return false
	}
	if users == nil {
		return false
	}
	for index, _ := range users {
		if users[index].Post == "false" {
			return false
		}
	}
	return true
}

func adminCheck(username string, password string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	bytes := hash.Sum(nil)
	sum := hex.EncodeToString(bytes)
	var users []user
	sql := "SELECT id, username,password,allowadd,allowpost,admin,email FROM blog_user WHERE username=? AND password=? "
	err := db.Select(&users, sql, username, sum)
	if err != nil {
		fmt.Printf("user select failed   :", err, "\n")
		return false
	}
	if users == nil {
		return false
	}
	for index, _ := range users {
		if users[index].Admin != "true" {
			return false
		}
	}
	return true
}

func adminCheckById(userid int) bool {
	var users []user
	sql := "SELECT id, username,password,allowadd,allowpost,admin,email FROM blog_user WHERE id=?  "
	err := db.Select(&users, sql, userid)
	if err != nil {
		fmt.Printf("user select failed   :", err, "\n")
		return false
	}
	if users == nil {
		return false
	}
	for index, _ := range users {
		if users[index].Admin != "true" {
			return false
		}
	}
	return true
}

func getUserId(username string, password string) (int, bool) {
	var id int
	hash := sha256.New()
	hash.Write([]byte(password))
	bytes := hash.Sum(nil)
	sum := hex.EncodeToString(bytes)
	var users []user
	sql := "SELECT id, username,password,allowadd,allowpost,admin,email FROM blog_user WHERE username=? AND password=? "
	err := db.Select(&users, sql, username, sum)
	if err != nil {
		fmt.Printf("user select failed  :", err, "\n")
		return -1, false
	}
	if users == nil {
		return -1, false
	}
	for index, _ := range users {
		if users[index].Username == username {
			id = users[index].Id
			break
		}
	}
	return id, true
}

func addPage(username string, password string, title string, main string) bool {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return false
	}
	if addAllow(username, password) {
		userId, theBool := getUserId(username, password)
		if theBool {
			sql := "INSERT INTO blog_page(id,user_id,title,main,time) VALUES(DEFAULT,?,?,?,?)"
			r, err := db.Exec(sql, userId, title, main, time.Now().Unix())
			if err != nil {
				fmt.Println("exec failed   :", err)
				conn.Rollback()
				return false
			}
			_, err = r.LastInsertId()
			if err != nil {
				fmt.Println("exec failed   :", err)
				conn.Rollback()
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
	conn.Commit()
	return true
}

func getPagePosterId(pageid int) int {
	var pages []page
	sql := "SELECT id,user_id,title,main,time FROM blog_page WHERE id=?"
	err := db.Select(&pages, sql, pageid)
	if err != nil {
		fmt.Printf("page select failed   :", err, "\n")
		return -1
	}
	if pages == nil {
		return -1
	}
	return pages[0].UserId
}

func delPage(username string, password string, pageid int) bool {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return false
	}
	userid, _ := getUserId(username, password)
	if userid == getPagePosterId(pageid) || adminCheck(username, password) {
		sql := "DELETE FROM blog_page where id=?"
		res, err := db.Exec(sql, pageid)
		if err != nil {
			fmt.Println("exce failed   :", err)
			conn.Rollback()
			return false
		}
		_, err = res.RowsAffected()
		if err != nil {
			fmt.Println("row failed   :", err)
		}
	} else {
		return false
	}
	conn.Commit()
	return true
}

func editPage(username string, password string, pageid int, main string) bool {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return false
	}
	userid, _ := getUserId(username, password)
	if userid == getPagePosterId(pageid) || adminCheck(username, password) {
		sql := "UPDATE blog_page SET main=? WHERE id=?"
		res, err := db.Exec(sql, main, pageid)
		if err != nil {
			fmt.Println("exec failed   :", err)
			conn.Rollback()
			return false
		}
		_, err = res.RowsAffected()
		if err != nil {
			fmt.Println("rows failed", err)
		}
	} else {
		return false
	}

	return true
}

func getUsername(userid int) string { //返回值：该人的名字以及是否为管理员
	var users []user
	sql := "SELECT id, username,password,allowadd,allowpost,admin,email FROM blog_user WHERE id=?"
	err := db.Select(&users, sql, userid)
	if err != nil {
		fmt.Printf("page select failed   :", err, "\n")
		return "none"
	}
	if users == nil {
		return "none"
	}
	return users[0].Username
}

func getPageList(leastId int, mostId int) ([]int, []string, []string, []int) {
	var pages []page
	var idList []int
	var nameList []string
	var posterList []string
	var time []int
	sql := "SELECT id,user_id,title,main,time FROM blog_page WHERE ?<id AND id<?"
	err := db.Select(&pages, sql, leastId, mostId)
	if err != nil {
		fmt.Printf("page select failed   :", err, "\n")
		idList = append(idList, -1)
		nameList = append(nameList, "none")
		posterList = append(posterList, "none")
		time = append(time, 114514)
		return idList, nameList, posterList, time
	}
	if pages == nil {
		idList = append(idList, -1)
		nameList = append(nameList, "none")
		posterList = append(posterList, "none")
		posterList = append(posterList, "none")
		return idList, nameList, posterList, time
	}
	for index, _ := range pages {
		idList = append(idList, pages[index].Id)
		nameList = append(nameList, pages[index].Title)
		posterList = append(posterList, getUsername(pages[index].UserId))
		time = append(time, pages[index].Time)

	}

	return idList, nameList, posterList, time
}

func getPage(pageid int) string {
	var pages []page
	sql := "SELECT id,user_id,title,main,time FROM blog_page WHERE id=?"
	err := db.Select(&pages, sql, pageid)
	if err != nil {
		fmt.Printf("page select failed   :", err, "\n")
		return "none"
	}
	if pages == nil {
		fmt.Printf("no thing\n")
		return "none"
	}
	return pages[0].Main
}

//以下是和token有关

func createToken(username string, password string) string {
	var users []user
	hash := sha256.New()
	hash.Write([]byte(password))
	bytes := hash.Sum(nil)
	sum := hex.EncodeToString(bytes)
	sql := "SELECT id, username,password,allowadd,allowpost,admin,email FROM blog_user WHERE username=? AND password=?"
	err := db.Select(&users, sql, username, sum)
	if err != nil {
		fmt.Printf("user select failed   :", err, "\n")
		return "none"
	}
	if users == nil {
		fmt.Printf("no thing\n")
		return "none"
	}
	noHashToken := username + password + string(time.Now().Unix())
	hash.Write([]byte(noHashToken))
	tokenBytes := hash.Sum(nil)
	token := hex.EncodeToString(tokenBytes)
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed  :", err, "\n")
		return "none"
	}
	sql = "INSERT INTO blog_token(id,user_id,main_token,time) VALUES(DEFAULT,?,?,?)"
	userId, _ := getUserId(username, password)
	r, err := db.Exec(sql, userId, token, time.Now().Unix())
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		conn.Rollback()
		return "none"
	}
	_, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed  :", err, "\n")
		conn.Rollback()
		return "none"
	}
	conn.Commit()
	return token
}

func getUserIdByToken(token string) (int, bool) { //userid,isadmin
	var tokens []tokenType
	var id int
	var admin bool
	sql := "SELECT id,user_id,main_token,time FROM blog_token WHERE main_token=?"
	err := db.Select(&tokens, sql, token)
	if err != nil {
		fmt.Printf("user select failed   :", err, "\n")
		return -1, false
	}
	if tokens == nil {
		fmt.Printf("no thing\n")
		return -1, false
	}
	for index, _ := range tokens {
		if tokens[index].MainToken == token {
			id = tokens[index].UserId
			if adminCheckById(id) {
				admin = true
			} else {
				admin = false
			}
			break
		}
	}
	return id, admin
}

func checkToken(token string) (int, bool) { //user id,authorize success or fail
	var tokens []tokenType
	var id int
	sql := "SELECT id,user_id,main_token,time FROM blog_token WHERE main_token=?"
	err := db.Select(&tokens, sql, token)
	if err != nil {
		fmt.Printf("user select failed   :", err, "\n")
		return -1, false
	}
	if tokens == nil {
		fmt.Printf("no thing\n")
		return -1, false
	}
	id = tokens[0].UserId
	return id, true
}

func addAllowByToken(token string) bool {
	var users []user
	id, result := checkToken(token)
	if result == false {
		return false
	}
	sql := "SELECT id, username,password,allowadd,allowpost,admin,email FROM blog_user WHERE id=?"
	err := db.Select(&users, sql, id)
	if err != nil {
		fmt.Printf("user select failed   :", err, "\n")
		return false
	}
	if users == nil {
		return false
	}
	for index, _ := range users {
		if users[index].Add == "false" {
			return false
		}
	}
	return true
}

func postAllowByToken(token string) bool {
	var users []user
	id, result := checkToken(token)
	if result == false {
		return false
	}
	sql := "SELECT id, username,password,allowadd,allowpost,admin,email FROM blog_user WHERE id=?"
	err := db.Select(&users, sql, id)
	if err != nil {
		fmt.Printf("user select failed   :", err, "\n")
		return false
	}
	if users == nil {
		return false
	}
	for index, _ := range users {
		if users[index].Post == "false" {
			return false
		}
	}
	return true
}

func adminCheckBytoken(token string) bool {
	_, admin := getUserIdByToken(token)
	return admin
}

func addPageByToken(token string, title string, main string) bool {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return false
	}
	if addAllowByToken(token) {
		userId, theBool := checkToken(token)
		if theBool {
			sql := "INSERT INTO blog_page(id,user_id,title,main,time) VALUES(DEFAULT,?,?,?,?)"
			r, err := db.Exec(sql, userId, title, main, time.Now().Unix())
			if err != nil {
				fmt.Println("exec failed   :", err)
				conn.Rollback()
				return false
			}
			_, err = r.LastInsertId()
			if err != nil {
				fmt.Println("exec failed   :", err)
				conn.Rollback()
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
	conn.Commit()
	return true
}

func delPageByToken(token string, pageid int) bool {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return false
	}
	userid, admin := getUserIdByToken(token)
	if userid == getPagePosterId(pageid) || admin {
		sql := "DELETE FROM blog_page WHERE id=?"
		res, err := db.Exec(sql, pageid)
		if err != nil {
			fmt.Println("exce failed   :", err)
			conn.Rollback()
			return false
		}
		_, err = res.RowsAffected()
		if err != nil {
			fmt.Println("row failed   :", err)
		}
	} else {
		return false
	}
	conn.Commit()
	return true
}

func editPageByToken(token string, pageid int, main string) bool {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return false
	}
	userid, admin := getUserIdByToken(token)
	if userid == getPagePosterId(pageid) || admin {
		sql := "UPDATE blog_page SET main=? WHERE id=?"
		res, err := db.Exec(sql, main, pageid)
		if err != nil {
			fmt.Println("exec failed   :", err)
			conn.Rollback()
			return false
		}
		_, err = res.RowsAffected()
		if err != nil {
			fmt.Println("rows failed", err)
		}
	} else {
		return false
	}
	return true
}

func delToken(token string) bool {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return false
	}
	sql := "DELETE FROM blog_token WHERE main_token=?"
	res, err := db.Exec(sql, token)
	if err != nil {
		fmt.Println("exce failed   :", err)
		conn.Rollback()
		return false
	}
	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println("row failed   :", err)
		return false
	}
	conn.Commit()
	return true
}
