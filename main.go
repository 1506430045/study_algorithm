package main

import "fmt"

type Student struct {
	Id int
	Name string
}

func main() {
	a := new(map[int]int)
	b := make(map[int]int, 6)

	fmt.Println(a, b)
}


