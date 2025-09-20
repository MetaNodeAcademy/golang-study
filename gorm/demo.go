package main

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct { //建表时表名自动添加‘s’
	gorm.Model
	Code  string
	Name  string `gorm:"size:260"`
	Price uint
	TagNo string //驼峰转蛇形  tag_no
}

func main() {
	dsn := "root:root@/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //使用默认配置
	if err != nil {
		panic("failed to connect database")
	}

	//准备上下文
	ctx := context.Background()

	fmt.Println("success to connect database")
	db.AutoMigrate(&Product{}) //自动创建表(若表不存在)，可多次执行，当有新增或变更字段时则同步更新

	//新增
	db.Create(&Product{
		Code:  "P01",
		Name:  "水杯",
		Price: 100,
		TagNo: "666",
	})
	db.Create(&Product{
		Code:  "P02",
		Name:  "水壶",
		Price: 200,
		TagNo: "888",
	})
	db.Create(&Product{
		Code:  "P04",
		Name:  "雨伞",
		Price: 400,
		TagNo: "000",
	})
	//查询
	product, err := gorm.G[Product](db).Where("code = ?", "P01").Find(ctx)
	fmt.Printf("查询结果1:%v,%v", product, err)

	product2, err := gorm.G[Product](db).Where("id = 1").First(ctx)
	fmt.Printf("查询结果2:%v,%v", product2, err)

	//更新

	var product3 Product
	result := db.Where("code = ?", "P01").First(&product3)
	if result.Error == nil {
		//更新
		db.Model(&Product{}).Where("code = ?", "P01").Updates(map[string]interface{}{"price": 345, "tag_no": "999"})
	}
	gorm.G[Product](db).Where("code = ?", "P02").Update(ctx, "price", 350)
	//删除
	// db.Where("code = ?", "P01").Delete(&Product{})
	// gorm.G[Product](db).Where("id = ?", product3.ID).Delete(ctx)
	gorm.G[Product](db).Where("code = ?", "P04").Delete(ctx)

	fmt.Println("基本操作执行完成...")

}
