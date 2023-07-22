package main

import (
	"fmt"
)

func updTitle(book Book) {
	book.title = "C"
}

func updTitleP(book *Book) {
	book.title = "C"
}

type Book struct {
	title string
	auth  string
}

func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "me"
	fmt.Println(book1)

	updTitle(book1)
	fmt.Println(book1)

	updTitleP(&book1)
	fmt.Println(book1)

	person1 := new(Person)
	person1.SetName("King")
	person1.GetName()
	person1.SetName("Queen")
	person1.GetName()

	stu1 := new(Student)
	stu1.SetID(1)
	stu1.GetID()

	dog := &Dog{
		"Yellow",
	}

	cat := &Cat{
		"Black",
	}

	AnimalSleep(dog)
	AnimalSleep(cat)
}
