package main

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//hook钩子函数包括 BeforeCreate BeforeSave AfterSave AfterCreate

type Product struct {
	gorm.Model
	Code        string
	Name        string `gorm:"size:260"`
	Price       uint
	TagNo       string       //驼峰转蛇形  tag_no
	Alias       *string      //指针类型，表示允许空值 null
	ProductDate sql.NullTime //允许空值
}

func main() {

	dsn := "root:root@/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	context := context.Background()

	fmt.Printf("db=%v,context=%v", db, context)

}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Price < 5 {
		fmt.Println("价格不能低于5RMB")
	}
	fmt.Println("BeforeCreate 完成")
	return
}

func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("BeforeSave 完成")
	return
}

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate 完成")
	return
}

func (p *Product) AfterSave(tx *gorm.DB) (err error) {
	fmt.Println("AfterSave 完成")
	return
}
