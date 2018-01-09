package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Human interface {
	say()
	eat()
}

func (this *Person) say() {
	fmt.Printf("say:%+v\n", this)
	this.name = "change_it"
}

func (this *Person)  eat() {
	fmt.Printf("eat:%+v\n", this)
	this.name = "change_it_again"
}

func main() {
	var human Human = &Person{name:"123", age:1}
	human.say()
	human.eat()

	fmt.Printf("final val=%+v\n", human)
}
