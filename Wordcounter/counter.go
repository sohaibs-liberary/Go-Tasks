package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Wordcounter(content string) {
  
	word := strings.Fields(content)

	fmt.Printf("Total words: %d\n", len(word))
}

func main() {
	data, _ := os.Open("counter.txt")
	defer data.Close()

	r := bufio.NewReader(data)

	for {
		line, _, err := r.ReadLine()

		if len(line) > 0 {
			Wordcounter(string(line))
		}

		if err != nil {
			fmt.Printf("This is an error %v\n", err)
			break
		}
	}
}
