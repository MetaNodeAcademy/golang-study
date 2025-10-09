package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string `gorm:"column:name;type:varchar(255);not null"`
	Age   int    `gorm:"column:age;type:int(11);not null"`
	Grade string `gorm:"column:grade;type:varchar(255);not null"`
}

type Account struct {
	gorm.Model
	Balance float32 //要大写，要不然gorm无法访问这些字段，导致数据不能正常入库
}

type Transaction struct {
	gorm.Model
	FromAccountId int
	ToAccountId   int
	Amount        float32
}

func main() {
	dsn := "root:root@/gorm-test?charset=utf8&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Student{})
	fmt.Println("数据表创建完成")

	db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})
	//查询年龄大于18的学生
	var students []Student
	db.Where("age>=?", 18).Find(&students)
	for _, student := range students {
		fmt.Printf("学生姓名:%s,年龄:%d,年级:%s\n", student.Name, student.Age, student.Grade)
	}
	// db.Where("name=?", "张三").Update("grade", "四年级")
	db.Model(&Student{}).Debug().Where("name=?", "张三").Update("grade", "四年级")
	db.Delete(&Student{}, "age<?", 15)

	//////////////////////////////////////////////////////////////////////
	db.AutoMigrate(&Account{}, &Transaction{})
	db.Create(&Account{Balance: 1000})
	db.Create(&Account{Balance: 500})
	//开始事务
	db.Begin()
	accountA := Account{}
	db.Where("id=?", 1).Find(&accountA)
	if accountA.Balance < 100 {
		db.Rollback()
		return
	}
	//变更A账户
	db.Model(&Account{}).Where("id=?", 1).Update("balance", accountA.Balance-100)
	//变更B账户
	accountB := Account{}
	db.Where("id=?", 2).Find(&accountB)
	db.Model(&Account{}).Where("id=?", 2).Update("balance", accountB.Balance+100)
	//写Transaction记录
	db.Create(&Transaction{FromAccountId: 1, ToAccountId: 2, Amount: 100})
	//提交事务
	db.Commit()
	fmt.Println("转账成功")
}
