package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID     int
	Name   string
	Code   string
	Emp    *[]string `gorm:"type:text"`
	NameEn string
}

func main() {

	dsn := "root:root@/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //使用默认配置
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	fmt.Println("数据表创建完成")
}
