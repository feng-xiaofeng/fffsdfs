package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //数据池
var err error

func init() {
	dsn := "root:mm2580AA@tcp(127.0.0.1:3306)/my_demo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	////数据库自动创建表 自动迁移（把结构体和数据表进行对呀）
	DB.AutoMigrate(&User{})
}
