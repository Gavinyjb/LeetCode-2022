package main

import "fmt"

type Person interface {
	Greet()
}
type person struct {
	Name string
	Age  int
}

func (p person) Greet() {
	fmt.Printf("你好 我的名字是%s\n", p.Name)
}
func NewPerson(name string, age int) Person {
	return person{
		Name: name,
		Age:  age,
	}
}

func main() {
	var p Person
	p = NewPerson("yvjinbo", 23)
	fmt.Println(p)
	p.Greet()
}
