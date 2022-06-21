package main

import (
	"fmt"
	"log"
	"os"
)

var data []int = []int{1, 2, 3, 4, 5}

func writeToChan(writeChan chan<- int) {
	defer close(writeChan)
	for _, v := range data {
		writeChan <- v
	}
}

func chanOwner() <- chan int {
	results := make(chan int, 5)

	go func() {
		defer close(results)

		for i := 1; i <= 5; i++ {
			results <- i
		}
	}()

	return results
}

func consumer(results <-chan int) {
	for result := range results {
		fmt.Println(result)
	}
}

type Result struct {
	Response *os.File
	Error error
}

func checkFiles(done <-chan interface{}, filenames ...string) <-chan Result {
	results := make(chan Result)

	go func() {
		defer close(results)
		
		for _, filename := range filenames {
			var result Result
			file, err := os.Open(filename)
			result = Result{file, err}
			// if err != nil {
			// 	log.Println(err)
			// 	return
			// }

			select {
			case <-done:
				return
			case results <- result:
			}
		}
	}()

	return results
}



func main() {
	

	// done := make(chan interface{})

	// defer close(done)

	// filenames := []string{"main.go", "x.go", "y.go"}

	// for result := range checkFiles(done, filenames...) {
	// 	if result.Error != nil {
	// 		log.Printf("error: %v\n", result.Error)
	// 		continue
	// 	}
		
	// 	fmt.Printf("Respose: %v\n", result.Response)
	// }

	// results := chanOwner()

	// consumer(results)

	// handleData := make(chan int)

	// go writeToChan(handleData)

	// for integer := range handleData {
	// 	fmt.Println(integer)
	// }
}