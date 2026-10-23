package main

import "fmt"

func main() {
	height_map := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}
	for len(height_map) != 0 {
		tallest := ""
		height := 0
		for key, value := range height_map {
			if value > height {
				tallest = key
				height = value
			}
		}
		fmt.Print(tallest, " ")
		delete(height_map, tallest)
	}
}
