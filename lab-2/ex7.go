package main

import (
	"fmt"
	"time"
)

type Coordinates struct {
	x int
	y int
}

func robot(name string, path []Coordinates, channel chan<- Coordinates, other_robot <-chan Coordinates, report chan<- string) {
	for _, pos := range path {
		channel <- pos
		time.Sleep(100 * time.Millisecond)
	}
	close(channel)

	for pos := range other_robot {
		fmt.Println(name, "a primit { x:", pos.x, "y:", pos.y, "}")
	}

	report <- name + " gata"
}

func main() {
	canal1 := make(chan Coordinates, 10)
	canal2 := make(chan Coordinates, 10)
	report := make(chan string)

	go robot("Robot 1", []Coordinates{{1, 1}, {2, 2}, {3, 3}}, canal1, canal2, report)
	go robot("Robot 2", []Coordinates{{10, 10}, {11, 11}}, canal2, canal1, report)

	fmt.Println(<-report)
	fmt.Println(<-report)
	fmt.Println("S-au terminat robotii")
}
