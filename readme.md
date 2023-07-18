# Benzenz-Blog
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
        - 安装go与mysSQL        - ~~安装依赖beego~~ 貌似不需要
        - 安装go第三方库
        ```
        go get go-blog-microserver
        ```
        - 初始化MySQL
            打开controllers/sql.go，在31行：
            ```golang
            database, err := sqlx.Open("mysql", "bloguser:blogpwd@tcp(127.0.0.1:3306)/blogdb")
            //                                  用户名    数据库密码     主机及其端口    数据库名称
            ```
            进入MySQL，把blogdb.sql导入
        -运行main.go

__具体的后端请求请在[这里](doc/help.md)中查看__