package initialize

import (
	"fmt"
	"github.com/betterego/go-better-admin/server/global"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"

)

func init() {
	global.DB = initMysql()
}

func initMysql() *gorm.DB {
	mysql := global.CONFIG.Mysql
	if mysql.Dbname == "" {
		return nil
	}
	dsn := mysql.Dsn()
	mysqlConfig :=driver.Config{
		DSN: dsn, // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	if db, err := gorm.Open(driver.New(mysqlConfig), &gorm.Config{}); err != nil {
		fmt.Println("数据库链接异常！")
		return nil
	}else {
		sqlDB,_ := db.DB()
		sqlDB.SetMaxIdleConns(mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(mysql.MaxOpenConns)
		return db
	}

}