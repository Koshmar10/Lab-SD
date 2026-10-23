package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	array := []int{2, 7, 11, 8, 4, 9}
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	sum, err := strconv.Atoi(scanner.Text())

	if err != nil {
		println("Provide a valid integer")
		return
	}

	seen_index := map[int]int{}
	for index, elem := range array {
		needed := sum - elem
		other_index, found := seen_index[needed]
		if found {
			fmt.Println(other_index, index)
		}
		seen_index[elem] = index
	}
}
