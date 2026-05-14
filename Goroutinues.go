package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type result struct {
	words      int
	letters    int
	lines      int
	spaces     int
	vowels     int
	paragraphs int
	special    int
}

// chunk  data

func countChunk(chunk []rune) result {

	var r result

	inWord := false

	for _, ch := range chunk {

		switch ch {
		case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '~', '`':
			r.special++

		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			r.vowels++
		}
		switch ch {
		case ' ':
			r.spaces++
			inWord = false
		case '\t':
			inWord = false
		case '\n':
			r.lines++
			inWord = false
		default:
			if !inWord {
				r.words++
				inWord = true
			}
		}

	}
	return r
}

func main() {
	start := time.Now()

	fileData, _ := os.ReadFile("word.txt")

	runes := []rune(string(fileData))
	total := len(runes)

	// Number of Goroutinues

	numWorkers := 2
	chunkSize := (total + numWorkers - 1) / numWorkers

	resultCh := make(chan result, numWorkers)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {

		start := i * chunkSize
		end := start + chunkSize

		if end > total {

			end = total

		}
		if start >= total {

			break

		}
		wg.Add(1)
		go func(chunk []rune) {
			defer wg.Done()
			resultCh <- countChunk(chunk)
		}(runes[start:end])
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var final result
	for r := range resultCh {
		final.words += r.words
		final.letters += r.letters
		final.lines += r.lines
		final.special += r.special
		final.spaces += r.spaces
		final.paragraphs += r.paragraphs
		final.vowels += r.vowels
	}

	fmt.Printf("Total Words: %d\n", final.words)
	fmt.Printf("Total Letters: %d\n", final.letters)
	fmt.Printf("Total Lines: %d\n", final.lines)
	fmt.Printf("Total Special: %d\n", final.special)
	fmt.Printf("Total Spaces: %d\n", final.spaces)
	fmt.Printf("Total Paragraphs: %d\n", final.paragraphs)
	fmt.Printf("Total Vowels:%d\n", final.vowels)
	fmt.Println("Execution time:", time.Since(start))
}






















// for _, ch := range string(fileData) {
// 	letters++

// 	switch ch {
// 	case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '~', '`':
// 		special++
// 	}

// 	if ch == ' ' {
// 		spaces++
// 	}
// 	if ch == '\n' {
// 		lines++
// 	}

// 	if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
// 		inWord = false
// 	} else if !inWord {
// 		words++
// 		inWord = true
// 	}

// }

// 	fmt.Printf("Total Words %d \n", words)
// 	fmt.Printf("Total Letter %d \n"  , letters)
// 	fmt.Printf("Total Lines %d \n"  ,lines)
// 	fmt.Printf("Total Special %d \n"  ,special)
// 	fmt.Printf("Total Lines %d \n"  ,lines)
// 	fmt.Printf("Total Paragraphes %d \n"  ,paragraphes)
