package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	just_even := []int{}
	
	fmt.Println("Toate", array)
	
	for _, elem := range array {
		if elem%2 == 0 {
			just_even = append(just_even, elem)
		}
	}

	fmt.Println("Doar pare ", just_even)
}
