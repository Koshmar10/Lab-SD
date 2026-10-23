package main

import (
	"fmt"
)

func main() {
	unsorted_array := []int{49, 100, 20, 510, 40, 1, 2, 3}
	fmt.Println("Unsorted array ", unsorted_array)

	for i := 0; i < len(unsorted_array); i++ {
		for j := i + 1; j < len(unsorted_array); j++ {
			if unsorted_array[j] < unsorted_array[i] {
				aux := unsorted_array[i]
				unsorted_array[i] = unsorted_array[j]
				unsorted_array[j] = aux
			}
		}
	}
	
	fmt.Println("Sorted array ", unsorted_array)
}
