# Benzene-Blog
__BY TrafficLight6__  
![LOGO](./doc_img/logo.png)  
~~一个不怎么微服务的微服务项目~~  
此仓库为后端文件  
数据库为MySQL，使用了事务  
基于beego框架开发，api返回格式均为json
## 下载与部署
- 下载
    用git clone或wget下载，再不行可以有scp协议上传目标服务器
-部署
    - 直接使用编译好的Windows/Linux可执行文件（不推荐，因为sql连接是在controllers/sql.go文件内的）
        ~~部署个蛋都编译好了~~
    - 使用源代码部署
        - 安装go与mysSQL
        - ~~安装依赖beego~~ 貌似不需要
        - 安装go第三方库
        ```
        go get benzene-blog
        ```
        - 初始化MySQL
            打开controllers/sql.go，在31行：
            ```go
            database, err := sqlx.Open("mysql", "bloguser:blogpwd@tcp(127.0.0.1:3306)/blogdb")
            //                                  用户名    数据库密码     主机及其端口    数据库名称
            ```
            进入MySQL，把blogdb.sql导入
        - 配置邮箱验证进入controllers/email.go，从48行开始，注释处有*！*的需要修改：
            ```go
                server := mail.NewSMTPClient()
                server.Host = "smtp.host.com"   //你的邮箱的smtp服务器地址*！*
                server.Port = 114514    //你的邮箱的smtp服务器端口*！*
                server.Username = "email add"   //你的邮箱*！*
                server.Password = "smtp code"   //你的邮箱的smtp密码或授权码*！*
                server.Encryption = mail.EncryptionTLS

                smtpClient, err := server.Connect()
                if err != nil {
                    fmt.Println(err)
                    conn.Rollback()
                    return false
                }

                // Create email
                email := mail.NewMSG()
                email.SetFrom("")//你在邮件中的自称可自行修改
                email.AddTo(emailadd)
                email.SetSubject("Email Code")

                htmlStr := "<h1>your code is " + strconv.Itoa(code) + "</h1>"//邮箱信息可自行修改
                email.SetBody(mail.TextHTML, htmlStr) //发送html信息可自行修改
                // Send email
                err = email.Send(smtpClient)
                if err != nil {
                    fmt.Println(err)
                    conn.Rollback()
                    return false
                }
                conn.Commit()
                return true
            ```

        - 禁用冗余服务
            - 自行在router/router.go内将不需要的服务条目注释掉
        - 运行main.go
- 注意，用token验证身份（具体请参照doc/help.md）的需要在前端cookie储存返回的token

## 更新日志
    - beta 0.1.0 2023年7月18日
    - 第一个版本 
    - beta 0.1.1 2023年7月22日
        - 关闭了emailCheck函数的服务
        - signup函数的服务现在要邮箱验证码
        - 增加了修改密码服务
    - beta 0.1.2 2023年7月23日
        - 增加了changePasswordByToken和getUserIdByToken的服务
        - 优化了controllers目录下的文件夹结构
        - 修改了部分路由请求

__具体的后端请求请在[这里](doc/help.md)中查看__