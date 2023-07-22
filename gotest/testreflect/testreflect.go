package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string `info:"name" doc:"my name"`
}

func tag(str interface{}) {
	// Elem()获取指针指向的对象本身
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, "doc: ", tagdoc)
	}
}

func testReflect(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("Value: ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.234

	testReflect(num)

	var p person

	tag(&p)
}
