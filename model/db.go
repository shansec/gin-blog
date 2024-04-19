package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	dns := "root:root@tcp(127.0.0.1:3306)/myginblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("连接数据库错误,检查参数", err)
	}

	db.AutoMigrate(&User{}, &Article{}, &Category{}, &Profile{})
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(10 * time.Second)

}
