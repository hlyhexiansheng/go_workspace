package main

import "fmt"

type Person struct {
	name string
	age  int
}

type human interface {
	say()
	eat()
}

func (p *Person) say()  {
	fmt.Printf("say:%+v\n",p)
	p.name = "fadsas"
}

func main() {
	p1 := Person{name:"123", age:1}

	fmt.Println(p1)
	//fmt.Println(*p1)
	p1.say()
	fmt.Println(p1)

}
