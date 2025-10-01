package main

import "fmt"

type Shape interface {
	Area() float64
	Perimter() float64
}

type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
func (r *Rectangle) Perimter() float64 {
	return 2 * (r.height + r.width)
}
func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c *Circle) Perimter() float64 {
	return 2 * 3.14 * c.radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person
	EmployeeId int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name:%s,Age:%d,EmployeeId:%d", e.Person.Name, e.Person.Age, e.EmployeeId)
}

func main() {
	r := Rectangle{width: 10, height: 20}
	c := Circle{radius: 5}
	fmt.Printf("矩形的面积是：%f\n", r.Area())
	fmt.Printf("矩形的周长是：%f\n", r.Perimter())
	fmt.Printf("圆形的面积是：%f\n", c.Area())
	fmt.Printf("圆形的周长是：%f\n", c.Perimter())

	e := Employee{Person: Person{Name: "张三", Age: 18}, EmployeeId: 1001}
	e.PrintInfo()
}
