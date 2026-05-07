package main

import (
	"bufio"
	"fmt"
	"os"
	

)

func countWords(content string) int {


	count := 0
	inWord := false

	for _, ch := range content {
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			inWord = false
		} else {
			if !inWord {
				count++
				inWord = true
			}
		}
	}
	return count


}

func main() {
	data, _ := os.Open("counter.txt")
	defer data.Close()

  total := 0

  scanner := bufio.NewScanner(data)

  for scanner.Scan(){
    total +=countWords(scanner.Text())
  }
 fmt.Printf("Total Word : %d\n", total)
}

