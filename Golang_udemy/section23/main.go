package main

import (
	"fmt"
	// "log"
	"os"
	"time"
	"math/rand"
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

func DoSomething(done chan interface{}) <-chan int {
	readStream := make(chan int)

	go func() {
		defer fmt.Println("DoSomething Done")
		defer close(readStream)

		for {
			select {
			case readStream <- rand.Intn(100):
			case <-done:
				return
			}
		}
	}()

	return readStream

}

// func DoSomething(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	// completed := make(chan interface{})

	// go func() {
	// 	defer fmt.Println("DoSomething Done")
	// 	defer close(completed)

	// 	for {
	// 		select {
	// 		case s := <-strings:
	// 			fmt.Println(s)
	// 		case <-done:
	// 			return
	// 		}
	// 	}
	// 	// for s := range strings {
	// 	// 	fmt.Println(s)
	// 	// }
	// }()

	// return completed
// }

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})

	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:

			case <-or(append(channels[3:], orDone)...):
			}
		}

	}()

	return orDone

}

func signal(after time.Duration) <-chan interface{} {
	done := make(chan interface{})

	go func() {
		defer close(done)

		time.Sleep(after)
	}()

	return done

}

// func double(sl []int) []int {
// 	doubleSlice := make([]int, len(sl))

// 	for i, v := range sl {
// 		doubleSlice[i] = v * 2
// 	}
// 	return doubleSlice
// }

// func add(sl []int) []int {
// 	addSlice := make([]int, len(sl))

// 	for i, v := range sl {
// 		addSlice[i] = v + 1
// 	}

// 	return addSlice
// }

func generator(done <-chan interface{}, integer ...int) <-chan int {
	intStream :=make(chan int)

	go func() {
		defer close(intStream)

		for _, i := range integer {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()

	return intStream

}

func double(done <-chan interface{}, intStream <-chan int) <-chan int {
	doubleStream := make(chan int)

	go func() {
		defer close(doubleStream)

		for i := range intStream {
			select {
			case <-done:
				return
			case doubleStream <- i * 2:		
			}
		}
	}()

	return doubleStream

}

func add(done <-chan interface{}, intStream <-chan int) <-chan int {
	addStream := make(chan int)

	go func() {
		defer close(addStream)

		for i := range intStream {
			select {
			case <-done:
				return
			case addStream <- i + 1:		
			}
		}
	}()

	return addStream

}

func main() {
	

	// done := make(chan interface{})
	// defer close(done)

	// intStream := generator(done, 1, 2, 3, 4, 5)

	// for v := range double(done, double(done, add(done, intStream))) {
	// 	fmt.Println(v)
	// }

	// ints := []int{1,2,3,4,5}

	// for _, v := range double(add(ints)) {
	// 	fmt.Println(v)
	// }

	// start := time.Now()
	// <-or(signal(time.Hour), signal(time.Minute), signal(time.Second))
	// fmt.Printf("done after: %v\n", time.Since(start)) 

	// done1 := make(chan interface{})
	// done2 := make(chan interface{})
	// done3 := make(chan interface{})

	// done := make(chan interface{})
	// readStream := DoSomething(done)

	// for i := 1; i <= 3; i++ {
	// 	fmt.Println(<-readStream)
	// }
	
	// close(done)

	// time.Sleep(1 * time.Second)

	// fmt.Println("main Done")

	// done := make(chan interface{})
	// completed := DoSomething(done, nil)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	close(done)
	// }()

	// <-completed

	// fmt.Println("main Done")

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