package testfor

import "fmt"

func TestFor() {
	var array1 [10]int
	array2 := [10]int{1, 2, 3, 4}
	// slice
	// array3 := []int{1, 2, 3, 4}

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	for index, value := range array2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

}
