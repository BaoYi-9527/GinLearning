## GinLearning

### Reference

+ [《跟煎鱼学 Go》](https://eddycjy.com/go-categories/)
+ [Go ini 配置文件操作库](https://ini.unknwon.io/)
+ [Gin：Golang 的一个微框架，性能极佳](https://github.com/gin-gonic/gin)
+ [beego-validation：beego 的表单验证库，中文文档](https://github.com/astaxie/beego/tree/master/validation)
+ [gorm，对开发人员友好的 ORM 框架](https://gorm.io/zh_CN/docs/)
+ [com，一个小而美的工具包](https://github.com/Unknwon/com)
+ []()

### 1.起步

#### 1.1 项目初始化

**初始化项目并安装 Gin**

```bash
# 创建一个项目并进入该项目根目录
mkdir GinLearning && cd GinLearning
# 初始化 
go mod init GinLearning
# 获取 Gin 
go get -u github.com/gin-gonic/gin
# 安装 ini 库
go get -u gopkg.in/ini.v1
# 拉取 com 依赖包
go get -u github.com/unknwon/com
# 拉取 gorm 依赖包
go get -u github.com/jinzhu/gorm
# 拉取 MySQL 驱动
go get -u github.com/go-sql-driver/mysql
# beego-validation
go get github.com/astaxie/beego/validation
# jwt-go依赖
go get -u github.com/dgrijalva/jwt-go
```

#### 1.2 目录结构

```text
GinLearning/
├── conf            // 用于存储配置文件
├── middleware      // 应用中间件
├── models          // 应用数据库模型
├── pkg             // 第三方包
├── routers         // 路由逻辑处理
├── tmp             // 临时文件
├── resource        // 资源文件
└── runtime         // 应用运行时数据
```

```text
GinLearning/
├── conf
│   └── app.ini
├── main.go
├── middleware
│   └── jwt
│       └── jwt.go
├── models
│   ├── article.go
│   ├── auth.go
│   ├── models.go
│   └── tag.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       ├── jwt.go
│       └── pagination.go
├── routers
│   ├── api
│   │   ├── auth.go
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
├── runtime
```