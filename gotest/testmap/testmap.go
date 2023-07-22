package main

import "fmt"

func main() {
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("Empty map.")
	}

	map1 = make(map[string]string, 10)

	map1["one"] = "A"
	map1["two"] = "B"
	map1["three"] = "C"

	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["one"] = "A"
	map2["two"] = "B"
	map2["three"] = "C"
	for k, v := range map2 {
		fmt.Println("key: ", k)
		fmt.Println("value: ", v)
	}

	// map3 := map[string]string {
	// 	"one": "A",
	// 	"two": "B",
	// 	"three": "C",
	// }
}
