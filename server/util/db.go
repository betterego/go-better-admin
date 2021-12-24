package util

import (
	"errors"
	"fmt"
	"github.com/betterego/go-better-admin/server/config"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LinkDB(mysql config.Mysql) (*gorm.DB,error) {
	if mysql.Dbname == "" {
		return nil,errors.New("数据库链接失败，数据库名不能为空")
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
		fmt.Println("数据库链接失败！")
		return nil,err
	}else {
		sqlDB,_ := db.DB()
		sqlDB.SetMaxIdleConns(mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(mysql.MaxOpenConns)
		fmt.Println("数据库链接成功")
		return db,nil
	}
}