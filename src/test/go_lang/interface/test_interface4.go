package main

import "fmt"

//测试子类覆盖父类的方法

type I1 interface {
	m1()
	m2()
}

type Parents struct {
	name string
}

type Children struct {
	Parents
}

func (this *Parents) m1() {
	fmt.Println("parent's method.--[m1]")
}

func (this *Parents) m2() {
	fmt.Println("parent's method.--[m2]")
}

func (this *Children) m1() {
	fmt.Println("childrem's method.")
}

func main() {
	child := &Children{}
	child.m1()
	child.m2()
}
