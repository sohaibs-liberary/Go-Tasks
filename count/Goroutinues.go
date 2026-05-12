package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func count() {

	fileData, _ := os.ReadFile("word.txt")

	words := 0
	spaces := 0
	lines := 0
	letters := 0
	special := 0
	paragraphes := 0
	valve := 0
	inWord := false

	for _, ch := range string(fileData) {
		letters++

		switch ch {
		case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '~', '`':
			special++
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			valve++
		}

		switch ch {
		case ' ':
			spaces++
			inWord = false
		case '\t':
			inWord = false
		case '\n':
			lines++
			inWord = false
	
		default:
			if !inWord {
				words++
				inWord = true
			}
		}

	}

	fmt.Printf("Total Words %d \n", words)
	fmt.Printf("Total Letter %d \n", letters)
	fmt.Printf("Total Lines %d \n", lines)
	fmt.Printf("Total Special %d \n", special)
	fmt.Printf("Total Spaces %d \n", spaces)
	fmt.Printf("Total Paragraphes %d \n", paragraphes)
	fmt.Printf("Total Valve %d \n", valve)

}

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		count()
	}()

	wg.Wait()
	fmt.Println("Execution time:", time.Since(start))
}
