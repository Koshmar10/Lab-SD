package main

import (
	"fmt"
)

func first_channel(c chan string) {
	c <- "Mesaj canal 1"
}

func second_channel(c chan string) {
	c <- "Mesaj canal 2"
}

func main() {
	message_channel := make(chan string)

	go first_channel(message_channel)
	go second_channel(message_channel)

	first_message := <-message_channel
	second_channel := <-message_channel

	fmt.Printf("Primit de pe canal: %s\n", first_message)
	fmt.Printf("Primit de pe canal: %s\n", second_channel)
}
