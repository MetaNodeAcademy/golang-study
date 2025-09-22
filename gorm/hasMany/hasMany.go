package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type User struct {
// 	gorm.Model
// 	CreditCards []CreditCard `gorm:"foreignKey:UserId"` //重写外键，使用UserRef作为外键
// }

// type CreditCard struct { //创建后的表 credit_cards
// 	gorm.Model
// 	Number string
// 	UserId uint
// }

type User struct {
	gorm.Model
	MemberNumber string
	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
}

type CreditCard struct {
	gorm.Model
	Number     string
	UserNumber string
}

func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("CreditCards").Find(&users).Error
	return users, err
}

func main() {
	dsn := "root:root@/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //使用默认配置
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &CreditCard{})

	fmt.Println("数据表创建完成")
	//添加用户及信用卡信息
	db.Create(&User{
		MemberNumber: "1",
		CreditCards: []CreditCard{
			{Number: "1"},
			{Number: "2"},
		},
	})

	users, _ := GetAll(db)
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Printf("用户%v的信息为：%v\n", u.MemberNumber, u.CreatedAt)
		for _, c := range u.CreditCards {
			fmt.Printf("用户%v的信用卡信息为：%v\n", u.MemberNumber, c.Number)
		}
	}
}
