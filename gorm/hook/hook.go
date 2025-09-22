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

	fmt.Printf("db=%v,context=%v\n", db, context)

	//新增一条数据
	result := db.Create(&Product{Code: "SNK147", Name: "水杯", Price: 100, TagNo: "666"})
	fmt.Printf("新增后的结果为：%v\n", result)

	//查询一条数据
	var myProduct Product
	db.First(&myProduct, "code=?", "SNK147")
	fmt.Printf("查询结果为：%v", myProduct)

	//更新
	//这种情况下gorm基于安全考虑不会更新任何数据
	// db.Debug().Model(&Product{}).Updates(map[string]interface{}{"price": 888, "tag_no": "888"})

	//可以通过拼接where条件来更新
	db.Debug().Model(&Product{}).Where("1=1").Updates(map[string]interface{}{"price": 777, "tag_no": "777"})

	// db.Debug().Model(&myProduct).Updates(map[string]interface{}{"price": 888, "tag_no": "888"})

	db.Delete(&Product{}, "id=?", 7)
	fmt.Printf("更新后的结果为：%v\n", myProduct)

}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	//顺序2
	if p.Price < 5 {
		fmt.Println("价格不能低于5RMB")
	}
	fmt.Println("BeforeCreate 完成")
	return
}

func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	//顺序1
	fmt.Println("BeforeSave 完成")
	return
}

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	//顺序3
	fmt.Println("AfterCreate 完成")
	return
}

func (p *Product) AfterSave(tx *gorm.DB) (err error) {
	//顺序4
	fmt.Println("AfterSave 完成")
	return
}
