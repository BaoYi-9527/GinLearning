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