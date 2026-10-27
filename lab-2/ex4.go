package main

import (
	"fmt"
	"time"
)

func proces(ok chan bool) {
	fmt.Println("A inceput procesarea")

	time.Sleep(3000 * time.Millisecond)

	fmt.Println("S-a terminat procesarea")
	ok <- true
}

func main() {
	rx := make(chan bool)
	go proces(rx)
	<-rx
	fmt.Println("Programul s-a terminat")
}