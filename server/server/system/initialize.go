package system

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/betterego/go-better-admin/server/config"
	"github.com/betterego/go-better-admin/server/domain/system/request"
	"github.com/betterego/go-better-admin/server/global"
	"github.com/betterego/go-better-admin/server/util"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type InitializeService struct {
}



func (initializeService *InitializeService) InitDb(dbConfig *request.DBConfig) error  {
	if dbConfig.Host == "" {
		dbConfig.Host = "127.0.0.1"
	}
	if dbConfig.Port == "" {
		dbConfig.Port = "3306"
	}
	if dbConfig.UserName == "" {
		dbConfig.UserName = "root"
	}
	if dbConfig.DBName == "" {
		dbConfig.DBName = "gba"
	}
	if err := initializeService.createDB(dbConfig); err != nil {
		return err
	}
	mysqlConfig := *initializeService.getMysqlConfig(dbConfig)
	db,err := LinkDB(mysqlConfig)
	if err != nil {
		return err
	}
	global.DB = db

	// TODO 创建表

	// TODO 初始化表
	return nil
}

func (initializeService *InitializeService) WriteConfig(dbConfig *request.DBConfig) error {
	global.SYSTEM.Mysql = *initializeService.getMysqlConfig(dbConfig)
	cs := util.StructToMap(global.SYSTEM)
	for k, v := range cs {
		global.VIPER.Set(k, v)
	}
	return global.VIPER.WriteConfig()
}

func (initializeService *InitializeService) getMysqlConfig(dbConfig *request.DBConfig) *config.Mysql {
	path := fmt.Sprintf("%s:%s", dbConfig.Host, dbConfig.Port)
	return &config.Mysql{
		Path:     path,
		Dbname:   dbConfig.DBName,
		Username: dbConfig.UserName,
		Password: dbConfig.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}
}

func (initializeService *InitializeService) createDB(dbConfig *request.DBConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port)
	createDBSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", dbConfig.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createDBSql)
	return err
}

func LinkDB(mysql config.Mysql) (*gorm.DB,error) {
	if mysql.Dbname == "" {
		global.LOG.Error("数据库链接失败，数据库名不能为空")
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
		global.LOG.Error("数据库链接失败")
		return nil,err
	}else {
		sqlDB,_ := db.DB()
		sqlDB.SetMaxIdleConns(mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(mysql.MaxOpenConns)
		global.LOG.Info("数据库链接成功")
		return db,nil
	}
}