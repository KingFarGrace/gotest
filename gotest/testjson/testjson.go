package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Price  int    `json:"price"`
	Actors string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, "zhangbozhi"}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error: ", err)
	}

	fmt.Printf("jsonStr: %s\n", jsonStr)

	movie_ := Movie{}
	err = json.Unmarshal(jsonStr, &movie_)
	if err != nil {
		fmt.Println("json umarshal error: ", err)
	}

	fmt.Println(movie_)
}
