package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type Employee struct {
	Id         int64   `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type Book struct {
	Id     int64   `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	dsn := "root:root@/gorm-test?charset=utf8&parseTime=True&loc=Local"
	db, error := sqlx.Connect("mysql", dsn)
	if error != nil {
		panic(error)
	}
	defer db.Close()
	createTableSql := `CREATE TABLE IF NOT EXISTS employee (
		id INT AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		department VARCHAR(255) NOT NULL,
		salary DOUBLE NOT NULL,
		PRIMARY KEY (id)
	)`
	_, error = db.Exec(createTableSql)
	if error != nil {
		log.Fatalln("创建表失败")
	}
	// 插入数据
	INSERTSQL := `insert into employee(name,department,salary) values(?,?,?)`
	db.Exec(INSERTSQL, "奥沙利文", "技术部", 5000)
	db.Exec(INSERTSQL, "希金斯", "财务部", 3000)
	db.Exec(INSERTSQL, "特鲁姆普", "技术部", 4000)
	//查询数据
	SELECTSQL := `select * from employee where department=?`
	rows, error := db.Query(SELECTSQL, "技术部")
	if error != nil {
		log.Fatalln("查询数据失败")
	}
	defer rows.Close()
	var employeeList []Employee
	//获取部门为技术部的员工列表并封装到employeeList中
	db.Select(&employeeList, "select * from employee where department=?", "技术部")
	for emp := range employeeList {
		fmt.Printf("员工姓名:%s,员工编号:%d,员工部门:%s,员工工资:%f\n", employeeList[emp].Name, employeeList[emp].Id, employeeList[emp].Department, employeeList[emp].Salary)
	}
	var employee Employee
	err := db.Get(&employee, "select * from employee order by salary desc limit 1")
	if err != nil {
		fmt.Printf("查询失败:%v\n", err)
		return
	}
	fmt.Printf("工资最高的员工姓名:%s,员工编号:%d,员工部门:%s,员工工资:%f\n", employee.Name, employee.Id, employee.Department, employee.Salary)

	//////////////////////////////////////////////////////////////////////////////////
	crateTableBook := "CREATE TABLE IF NOT EXISTS book(\n" +
		"id int primary key auto_increment,\n" +
		"title varchar(255),\n" +
		"author varchar(255),\n" +
		"price float\n" +
		")"
	_, err1 := db.Exec(crateTableBook)
	if err1 != nil {
		fmt.Printf("创建表失败，具体错误: %v\n", err1)

		panic("failed to create table book")
	}
	//添加数据
	insertBook := `insert into book(title,author,price) values(?,?,?)`
	db.Exec(insertBook, "《西游记》", "吴承恩", 88.8)
	db.Exec(insertBook, "《红楼梦》", "曹雪芹", 99.9)
	db.Exec(insertBook, "《三国演义》", "罗贯中", 66.6)
	var bookes []Book
	er := db.Select(&bookes, "select * from book where price>?", 50)
	if er != nil {
		panic("failed to query book")
	}
	for _, book := range bookes {
		fmt.Printf("价格超过50元的书籍: %+v\n", book.Title)
	}
}
