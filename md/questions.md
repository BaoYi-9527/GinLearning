## 常见问题汇总

### 1. sql: unknown driver “mysql“ (forgotten import?)

在 `gorm` 中使用 `mysql` 驱动的时候，没有导入 `mysql` 驱动包。

```go
import (
	"GinLearning/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"	// 很重要!!! 引入 MySQL 驱动
	"log"
)
```

### 2. 代码变动未生效

Golang是一门编译型语言，与PHP、Python这些解释性语言不一样，需要编译后执行。

### 3. runtime error: invalid memory address or nil pointer dereference

```go
var db *gorm.DB
//...
func init()  {
	// ...
    db, err := gorm.Open(dbType, fmt.Sprintf(
        "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        user, password, host, dbName))
	// ...
}
```

虽然声明了全局变量 `db`，但是在 `init` 函数中，由于使用推导等号`:=`，`init` 函数中实际上生成了一个局部变量 `db` ,因此全局变量 `db` 并没有被赋值。

修改为:

```go
db, err = gorm.Open(dbType, fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName))
```

### 4. endless 在 windows 系统下安装失败

```bash
$ go get -u github.com/fvbock/endless
go: downloading github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
# github.com/fvbock/endless
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:64:3: undefined: syscall.SIGUSR1
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:65:3: undefined: syscall.SIGUSR2
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:68:3: undefined: syscall.SIGTSTP
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:111:5: undefined: syscall.SIGUSR1
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:112:5: undefined: syscall.SIGUSR2
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:115:5: undefined: syscall.SIGTSTP
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:119:5: undefined: syscall.SIGUSR1
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:120:5: undefined: syscall.SIGUSR2
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:123:5: undefined: syscall.SIGTSTP
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:224:3: undefined: syscall.Kill
D:\Go\pkg\mod\github.com\fvbock\endless@v0.0.0-20170109170031-447134032cb6\endless.go:224:3: too many errors
```

参考：*[windows下使用endless报错：undefined: syscall.SIGUSR1](https://blog.csdn.net/qq_28466271/article/details/116521955)*

1. `C:\Program Files\Go\src\syscall\types_windows.go` 下修改：

   ```go
   var signals = [...]string{
   	1:  "hangup",
   	2:  "interrupt",
   	3:  "quit",
   	4:  "illegal instruction",
   	5:  "trace/breakpoint trap",
   	6:  "aborted",
   	7:  "bus error",
   	8:  "floating point exception",
   	9:  "killed",
   	10: "user defined signal 1",
   	11: "segmentation fault",
   	12: "user defined signal 2",
   	13: "broken pipe",
   	14: "alarm clock",
   	15: "terminated",
   	/** 兼容windows start */
   	16: "SIGUSR1",
   	17: "SIGUSR2",
   	18: "SIGTSTP",
   	/** 兼容windows end */
   
   }
   
   /** 兼容windows start */
   func Kill(...interface{}) {
   	return;
   }
   const (
   	SIGUSR1 = Signal(0x10)
   	SIGUSR2 = Signal(0x11)
   	SIGTSTP = Signal(0x12)
   )
   /** 兼容windows end */
   ```

2. 可能会遇到 `Windows` 系统权限问题，无法修改源码；解决办法：

   + 鼠标右击 `types_windows.go` ，`属性->安全->编辑`  选择你的组或用户名后勾选相关权限后确认即可；
   + 如果找不到你的 `组或用户名`, `属性->安全->高级->添加->选择主体->高级->立即查找`

