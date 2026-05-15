package route

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type result struct {
    Words   int `json:"words"`
    Letters int `json:"letters"`
    Lines   int `json:"lines"`
    Spaces  int `json:"spaces"`
    Special int `json:"special"`

}

func Filecontext() result {

	var num int
	fileData, _ := os.ReadFile("word.txt")

	content := string(fileData)

	fmt.Println("Enter The Number of Goroutinues")
	fmt.Scan(&num)

	totalLen := len(content)

	worker := num
	chunkSize := (totalLen + worker - 1) / worker

	resultChan := make(chan result, worker)

	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 0; i < worker; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == worker-1 {
			end = totalLen
		}

		wg.Add(1)

		go func(chunk string) {
			defer wg.Done()

			var res result
			inWord := false

			for j := 0; j < len(chunk); j++ {
				ch := chunk[j]

				switch ch {
				case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=',
					'{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '~', '`':
					res.Special++
				}

				if ch == ' ' {
					res.Spaces++
				}
				if ch == '\n' {
					res.Lines++
				}

				if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
					inWord = false
				} else {
					if !inWord {
						res.Words++
						inWord = true
					}
					if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
						res.Letters++
					}
				}
			}

			resultChan <- res
		}(content[start:end])
	}

	// go func() {
	wg.Wait()
	close(resultChan)
	// }()

	finalWords := 0
	finalSpaces := 0
	finalLines := 0
	finalLetters := 0
	finalSpecial := 0

	for res := range resultChan {
		finalWords += res.Words
		finalSpaces += res.Spaces
		finalLines += res.Lines
		finalLetters += res.Letters
		finalSpecial += res.Special
	}

	fmt.Printf("Total Words:   %d\n", finalWords)
	fmt.Printf("Total Letters: %d\n", finalLetters)
	fmt.Printf("Total Lines:   %d\n", finalLines)
	fmt.Printf("Total Spaces:  %d\n", finalSpaces)
	fmt.Printf("Total Special: %d\n", finalSpecial)

	fmt.Println("Execution time:", time.Since(startTime))


    return result{
        Words:   finalWords,
        Letters: finalLetters,
        Lines:   finalLines,
       Spaces:  finalSpaces,
        Special: finalSpecial,
    }

}
