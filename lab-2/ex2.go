package main

import "fmt"

func send_msg(msg chan string) {
	recieved_msg := <-msg
	fmt.Println("Mesajul primit de la main: " + recieved_msg)
}

func main() {
	message_channel := make(chan string)

	go send_msg(message_channel)
	message_channel <- "Buna ce mai zici"

	fmt.Println("S-a terminat main()")
}
