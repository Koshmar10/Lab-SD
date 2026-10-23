package main

import (
	"fmt"
	"slices"
)


func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	fmt.Println("initial slice: ", slice)
	slices.Reverse(slice)
	fmt.Println("reversed slice: ", slice)

}