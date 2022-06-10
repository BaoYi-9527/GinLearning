package models

import (
	"GinLearning/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 很重要!!! 引入 MySQL 驱动
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"` // 软删除
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	//
	db, err = gorm.Open(dbType, fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	// 回调方法
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback 创建数据时会降序 CreateOn 和 ModifiedOn 字段的时间
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() { // 检查是否有错误
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok { // 判断是否有该字段
			if createTimeField.IsBlank { // 判断该字段值是否为空
				createTimeField.Set(nowTime) // 为空则给该字段赋值
			}
		}

		if modifyTimeField, ok := scope.FieldByName("modifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok { // 检查是否手动指定了 delete_option
			extraOption = fmt.Sprint(str)
		}
		deleteOnField, hasDeletedOnField := scope.FieldByName("DeleteOn") // 存在该字段则使用 UPDATE 软删除 否则使用 DELETE 硬删除
		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(), // 返回引用的表名
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(time.Now().Unix()),                  // 该方法可以添加值作为 SQL 的参数，也可防范 SQL 注入
				addExtraSpaceIfExists(scope.CombinedConditionSql()), // 返回组合好的SQL
				addExtraSpaceIfExists(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExists(scope.CombinedConditionSql()),
				addExtraSpaceIfExists(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExists(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
