package main

import "fmt"

type Animal interface {
	Sleep()
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (this Cat) Sleep() {
	fmt.Println("Cat is sleeping...")
}

func (this Dog) Sleep() {
	fmt.Println("Dog is sleeping...")
}

func AnimalSleep(animal Animal) {
	animal.Sleep()
}
