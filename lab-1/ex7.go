package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	word_count := map[string]int{}
	file, err := os.Open("ex7.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		for _, word := range words {
			word_count[word]++
		}
	}
	for key, value :=range word_count {
		fmt.Println(key,":", value, "times")
	}
}
