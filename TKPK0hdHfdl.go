package main

import "fmt"

func main() {
	factValue := 4
	arr := []int{}
	for x := 0; x < factValue; factValue-- {
		arr = append(arr, factValue)
	}
	fmt.Println(arr)

	total := soma(arr...)
	fmt.Println(total)
}

func soma(x ...int) int {

	initValue := 1

	for _, v := range x {
		initValue *= v
	}

	return initValue
}
