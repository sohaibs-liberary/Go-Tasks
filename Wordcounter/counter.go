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

// space counting

func spaceCounting(constant string) int {

	spaceCount := 0

	for _, ch := range constant {

		if ch == ' ' {
			spaceCount++
		}

	}
	return spaceCount

}

// Line Counting
func Linecounter(content string) int {

	Lineno := 0

	for _, ch := range content {

		if ch == '\n' {

			Lineno++
		}

	}
	return Lineno
}

// Letters
func Letterscounter(constant string) int {
	return len(constant)
}

func main() {
	data, _ := os.Open("counter.txt")
	defer data.Close()

	total := 0
	totalsapce := 0
	totalline := 0
	totalletter := 0

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		total += countWords(scanner.Text())
		totalsapce += spaceCounting(scanner.Text())
		totalline += Linecounter(scanner.Text())
		totalletter += Letterscounter(scanner.Text())

	}

	fmt.Printf("Total Word : %d\n", total)
	fmt.Printf("Total space : %d\n", totalsapce)
	fmt.Printf("Total Lines : %d\n", totalline)
	fmt.Printf("Total Letter : %d\n", totalletter)

}
