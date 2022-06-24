package main

import (
	"fmt"
	// "sync"
	// "Golang_udemy/myworkspace/person"
	"time"
	"os"
	"bufio"
	"strconv"
	"log"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func printGoodMorning() {
	fmt.Println(time.Now(), ": Good Morning")
}

func printHello() {
	fmt.Println(time.Now(), ": Hello")
}

func printGoodEvening() {
	fmt.Println(time.Now(), ": Good Evening")
}

func Say(ch1 chan int) {
	time.Sleep(10 * time.Second)
	x := <-ch1
	switch {
	case x == 1:
		printGoodMorning()
	case x == 2:
		printHello()
	case x == 3:
		printGoodEvening()
	default:
	}
}

func main() {
	// var wg sync.WaitGroup
	// message := 0
	ch1 := make(chan int)

	for {
		time.Sleep(2 * time.Second)
		go func() {
			fmt.Print("input number(1~4)?\n")
			// Scannerを使って一行読み
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			mem := scanner.Text()
			message, err := strconv.Atoi(mem)
			ch1 <- message
			if err != nil {
				log.Println(err)
			}
		}()
		
		go Say(ch1)
		go Say(ch1)
	}
	// fmt.Println("Hello World!")
	// fmt.Println(add(1,2))
	// fmt.Println(sub(1,2))


	// var personList person.PersonList
	// p1 := person.Person{"Yamada Taro", 24}
	// p2 := person.Person{"Bond", 55}
	// p3 := person.Person{"White", 72}
	// personList = append(personList, &p1, &p2, &p3)

	// fmt.Printf("Name: %v\nAge: %v\n", p1.Name, p1.Age)
	
	// wg.Add(2)
	// go person.UpdatePersonName(&wg, &p1, "Yamada Hanako")
	// go person.UpdatePersonAge(&wg, &p1, 22)
	// wg.Wait()

	// personList = person.RemovePerson(personList, 0)
	
	// fmt.Printf("Name: %v\nAge: %v\n", p1.Name, p1.Age)
	// fmt.Println(personList)
	// for _, v := range personList {
	// 	fmt.Println(*v)
	// }

}