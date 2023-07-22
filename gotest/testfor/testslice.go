package testfor

import "fmt"

func TestSlice() {

	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = make([]int, 3)
	slice2[0] = 10

	var slice3 []int = make([]int, 3)

	slice4 := make([]int, 3)
	slice4[0] = 100

	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice4: ", slice4)
}
