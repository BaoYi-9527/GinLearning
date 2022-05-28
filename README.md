## GinLearning

### Reference

+ [《跟煎鱼学 Go》](https://eddycjy.com/go-categories/)
+ [Go ini 配置文件操作库](https://ini.unknwon.io/)

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
