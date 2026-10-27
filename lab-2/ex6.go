package main

import (
	"fmt"
	"time"
)

func robot_race(name string, channel chan<- string) {
	if name == "Robot 1" {
		time.Sleep(2 * time.Second)
		channel <- "Robot 1 a terminat cursa"
	} else {
		time.Sleep(1 * time.Second)
		channel <- "Robot 2 a terminat cursa"
	}
}

func main() {
	robot1_finish := make(chan string)
	robot2_finish := make(chan string)

	go robot_race("Robot 1", robot1_finish)
	go robot_race("Robot 2", robot2_finish)

	select {
	case status := <-robot1_finish:
		fmt.Println(status)
	case status := <-robot2_finish:
		fmt.Println(status)
	}
}
