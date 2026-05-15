package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
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



func main() {

	time := time.Now()

	Filedata, _:= os.ReadFile("word.txt")
	content := string(Filedata)

	TotalLen := len(content)

	
	worker := 4
   chunkSize := (TotalLen + worker - 1) / worker

   returnCh := make(chan result, worker)

   // chunk data 

   for i := 0; i < worker; i++ {

	start := i*chunkSize
	end := chunkSize + start

	if i == worker -1 {

       end = TotalLen
		
	}
	wg.Add(1)

	go func(chunk) {

		defer wg.Done()

		returnCh <- result{}
		inWord := false

		for _, b := range chunk {
				ch := int(b)
				

				switch ch {
				case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '~', '`':
					res.special++
				}

				if ch == ' ' {
					res.spaces++
				}
				if ch == '\n' {
					res.lines++
				}

				if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
					inWord = false
				} else if !inWord {
					res.words++
					inWord = true
				}
			}

			resultChan <- res
		}(fileData[startIdx:endIdx])

		
	}()
	
   }



	

	fmt.Println("Execution Time", time.Since(time))

}
