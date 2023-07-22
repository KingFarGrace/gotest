package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func (this Person) GetName() {
	fmt.Println(this.name)
}

func (this *Person) SetName(name string) {
	this.name = name
}

type Student struct {
	Person
	id int
}

func (this Student) GetID() {
	fmt.Println(this.id)
}

func (this *Student) SetID(id int) {
	this.id = id
}

// func main() {
// 	var person1 = new(Person)
// 	person1.SetName("King")
// 	person1.GetName()
// 	person1.SetName("Queen")
// 	person1.GetName()
// }
