package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	gorm.Model
	Name      string
	Age       int
	CompanyId uint
	Company   Company
}

type Company struct {
	gorm.Model
	Name string
	Code string
}

//多对一
// func main() {

// 	dsn := "root:root@/gorm-test?charset=utf8&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //使用默认配置
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db.AutoMigrate(&User{})
// 	db.AutoMigrate(&Company{})
// 	fmt.Println("数据表创建完成")

// 	//方案1 先创建公司再创建用户
// 	// company:=Company{Name: "百度", Code: "baidu"}
// 	// db.Create(&company)
// 	// user := User{Name: "小王", Age: 18, CompanyId: company.ID}
// 	// db.Create(&user)
// 	//方案2 使用关联创建

// 	company2 := Company{Name: "阿里", Code: "alibaba"}
// 	user2 := User{Name: "小李", Age: 16, Company: company2}
// 	db.Create(&user2)

// }

//一对一

type Card struct {
	gorm.Model
	Number   string
	PersonId uint
}
type Person struct { //这里有个坑，创建数据库表时创建的表名为people,而不是persons(person的复数形式是people)
	gorm.Model
	Name string
	Card Card
}

// func main() {

// 	dsn := "root:root@/gorm-test?charset=utf8&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //使用默认配置
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db.AutoMigrate(
// 		&Person{},
// 		&Card{},
// 	)

// 	// card := Card{Number: "548989"}
// 	// person := Person{Name: "晓丽", Card: card}
// 	// db.Create(&person)

// 	var foundPerson Person
// 	result := db.Debug().Preload("Card").Find(&foundPerson)
// 	// result := db.Debug().Find(&foundPerson) //这里不preload Card，则打印的json中无card字段
// 	if result.Error != nil {
// 		panic(result.Error)
// 	} else {
// 		//将foundPerson打印成json
// 		json, _ := json.Marshal(foundPerson)
// 		fmt.Printf("查询结果为：%v\n", string(json))
// 	}

// 	//查询person表的所有数据
// 	var persons []Person
// 	result = db.Debug().Find(&persons)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	} else {
// 		json, _ := json.Marshal(persons)
// 		fmt.Printf("查询结果为：%v\n", string(json))
// 	}
// }

//多对多

type Employee struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:emp_lang;"` //多对多，当使用AutoMigrate创建Employee表时，会创建emp_lang关联表
}
type Language struct {
	gorm.Model
	LanName string
	starts  int
}

func main() {
	dsn := "root:root@/gorm-test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//表名不用复数形式
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(
		&Employee{},
		&Language{},
	)

	// employee := Employee{Name: "悟空", Languages: []Language{{LanName: "english", starts: 2}, {LanName: "中文", starts: 5}, {LanName: "法语", starts: 1}}}
	employee := Employee{Name: "唐僧", Languages: []Language{{LanName: "english"}}}

	db.Create(&employee) //会创建三张表 employee,language,emp_lang的相应数据
	var foundEmployee Employee
	db.Preload("Languages").Find(&foundEmployee)
	js, _ := json.Marshal(foundEmployee)
	fmt.Printf("查询结果为：%v\n", string(js))

	//查询同时掌握english的员工
	var employees []Employee
	db.Debug().Preload("language", "lan_name=?", "english").Find(&employees)
	js1, _ := json.Marshal(employees)
	fmt.Printf("查询结果为：%v\n", string(js1))

}
