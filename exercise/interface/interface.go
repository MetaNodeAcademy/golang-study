package main

import (
	"fmt"
)

type myInterface interface {
	playSnk(name string)
	listenMusic(name string)
}

type Person struct {
	name string
}

func (p Person) playSnk(name string) {
	fmt.Printf("当前name:%s正在玩斯诺克\n", p.name)
}

func (p Person) listenMusic(name string) {
	fmt.Printf("当前name:%s正在听音乐\n", p.name)
}

func (p Person) playOther(name string) {
	fmt.Printf("当前name:%s正在玩\n", p.name)
}

func (p Person) doSomething(inter myInterface) {
	fmt.Printf("%v开始执行doSomething方法\n", p.name)
}

func main() {
	p := Person{}
	p.name = "奥沙利文"
	var inter myInterface = p //这里，Person必须要实现myInterface中的所有方法，否则就会报错
	inter.playSnk(p.name)
	inter.listenMusic(p.name)

	p.playOther(p.name) //接口以外的方法，只能通过实例对象调用，而不能通过接口来调用

	var person2 = Person{
		name: "希金斯",
	}
	person2.doSomething(p)

}
