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

	go func(chunkSize) {

		defer wg.Done()

		returnCh <- result{}
		inWord := false

		
	}()
	
   }



	

	fmt.Println("Execution Time", time.Since(time))

}
