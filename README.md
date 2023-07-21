## Myblog_Server后端项目

这是一个具有多功能的个人博客网站。

🥤网站主要模块：博客、分类、标签、评论、附件、用户列表、我的信息、我的收藏、我的评论、热搜信息、友链、音乐、关于、系统公告、系统日志、系统信息。



🥤三大角色：用户、游客、管理员。游客可以看到部分后台的内容，但是没有关键数据的操作权限。



🥤支持一键开启/关闭全局评论，支持用户注册和忘记密码（提醒：前提是先绑定邮箱哦），支持夜间模式（有待优化）。



🥤用户绑定邮箱后，更新密码会通过邮箱提醒。

🥤支持七牛云上传。

🥤Jwt鉴权。

🥤调用[聚合数据API](https://dashboard.juhe.cn/)将用户IP转为物理地址。【每天50次免费使用，个人一般够用，每登录一次，算作一次使用机会】

🥤调用[天行数据API](https://www.tianapi.com/list/)获取热搜信息等。【每天100次免费使用，个人绝对够用。因为本项目是每30分调用一次Api，并将数据存储到mysql数据库，将数据通过接口传给前端。使用上不成问题。】







⛳️这是线上地址:  https://flowersbloom.com.cn.

⛳️网站演示账号：demo 密码：P@ssw0rd123..

⛳️默认管理员账号：admin 密码：admin



⛳️[Myblog_Web前端项目链接](https://github.com/flowersbloom18/myblog_web)



一些小bug正在修复中...

## 1. 后端技术栈

`go` `gin` `grom` `mysql` `redis`



## 2. 项目运行

```go 
// 安装环境
go mod tidy

go run main.go
```



## 3. 项目部署

### 3.1 以上传腾讯云服务器centos7为例，使用服务为宝塔面板

```yaml
# 交叉编译

# 在命令行里，这里记得要记住自己的GOARCH和GOOS的值
go env 

# 这里以Mac为例(本地运行的时候还要改回来)
export GOARCH=amd64
export GOOS=linux
# 🥤Windows的命令则为set

# 打包
go build -o main

# 如果不行再用这个
CGO_ENABLED=0 go build -o main

# 记得复原
export GOARCH="arm64"
export GOOS="darwin"
```

### 3.2 上传到服务器【前后端部署教程，简述版】

1、后端

```yaml
# 大前提【mysql、redis、yaml配置】

# 在mysql中创建myblog数据库，将myblog.sql数据库导入
# 连接redis
# 在本地配置好yaml文件

# 在适当的位置创建项目文件，
# 以本项目为例，先创建文件夹/www/wwwroot/myblog
前端：/www/wwwroot/myblog/myblog_web/dist
后端：/www/wwwroot/myblog/myblog_server/

# 上传exe、uploads、yaml文件到下
/www/wwwroot/myblog/myblog_server/

# 点击网站进入GO项目，项目启动配置。在8080端口
```

2、前端

```yaml
# 1.打包
npm run build -->dist. 

# 2.上传
dist-->/www/wwwroot/myblog/myblog_web/

# 3.创建并修改Nginx配置文件

# ================【大前提】==================
​	—>查看配置信息 nginx -t
​	—>修改配置信息 vi /www/server/nginx/conf/nginx.conf 
​	—>翻到底部，适当位置加入include /www/wwwroot/*/nginx_*.conf;


vi /www/wwwroot/myblog/nginx_myblog.conf

server {
    listen       80;
    server_name  yourdomain.com;


    location / {
      try_files $uri $uri/ /index.html;  # 解决刷新404问题
      root   /www/wwwroot/myblog/myblog_web/dist; # 位置
      index  index.html index.htm;
    }

    location /api/ { 
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header REMOTE-HOST $remote_addr;
      proxy_pass http://127.0.0.1:8080/api/; # 解决跨域问题
    }
    location /uploads/ {
      alias /www/wwwroot/myblog/myblog_server/uploads/; # 位置
    }
}

# 4.进入终端服务器平滑启动。
nginx -s reload
```

>  默认的管理员账号为：admin密码为：admin
>  网站演示的账号为：    demo 密码为：P@ssw0rd123..





至此，项目部署完毕。有问题可以到 https://flowersbloom.com.cn 讨论、交流。



