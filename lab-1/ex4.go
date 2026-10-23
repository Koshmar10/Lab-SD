package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	array := []float32{1.3, 2.0, 4.5, 10.2, 90.6}
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	number, err := strconv.Atoi(scanner.Text())

	if err != nil {
		println("Prvide a valid integer")
		return
	}

	if number >= len(array) {
		fmt.Println("index must be smaller than the length of the array")
		return
	}
	removed_element := array[number]
	array = append(array[:number], array[number+1:]...)
	fmt.Printf("Removed element %v, array: %v\n", removed_element, array)

}
