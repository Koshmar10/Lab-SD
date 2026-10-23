package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPrime(number int) bool {

	if number < 2 {
		return false
	}
	if number == 2 {
		return true
	} else if number%2 == 0 {
		return false
	}
	for i := 3; i <= number/2; i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	pare := [...]int{}
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	number, err := strconv.ParseInt(scanner.Text(), 10, 32)

	if err != nil {
		println("Prvide a valid integer")
		return
	}
	println(isPrime(int(number)))

	fmt.Println(pare)
}
