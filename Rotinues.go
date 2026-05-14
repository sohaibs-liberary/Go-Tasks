package main

import (
	"fmt"
	"os"
	"time"
)

type result struct{
	word int
	lineno int
	spaces int 
	vawles int 
	special int
	letters int 
}

func countWords(content string) int {

	// Wrod Counting

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

//No Of  Special Characters

func Spaceialchars(constant string) int {

	Special := 0

	for _, ch := range constant {
		
		switch ch {
		case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '~', '`'  :
			Special++
		}

	}
	return Special

}



func main() {

	time := time.Now()

	Filedata, _ := os.ReadFile("word.txt")
	contant := string(Filedata)

	for _, contant := range contant {


		
	}




	fmt.Println("Execution Time", time.Since(time))

}
