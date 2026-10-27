package main

import (
	"fmt"
)

func send_msg(msg chan string) {
	msg <- "salut"
	fmt.Println("Se trmite catre main")
}

func main() {
	ex_chan := make(chan string)

	go send_msg(ex_chan)

	msg := <-ex_chan
	fmt.Printf("Mesaj primit din alta goroutina: \"%s\"\n", msg)
}
