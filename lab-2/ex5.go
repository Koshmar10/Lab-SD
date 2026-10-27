package main

import "fmt"

func send(channel chan<- string, msg string) {
	channel <- msg
}

func receive(input <-chan string, output chan<- string) {
	msg := <-input
	output <- msg
}

func main() {
	send_channel := make(chan string)
	receive_channel := make(chan string)

	go send(send_channel, "salut")
	go receive(send_channel, receive_channel)

	fmt.Println(<-receive_channel)

}
