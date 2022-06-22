package main

import (
	"fmt"
	"sync"
	"Golang_udemy/myworkspace/person"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	var wg sync.WaitGroup

	// fmt.Println("Hello World!")
	// fmt.Println(add(1,2))
	// fmt.Println(sub(1,2))
	var personList person.PersonList
	p1 := person.Person{"Yamada Taro", 24}
	p2 := person.Person{"Bond", 55}
	p3 := person.Person{"White", 72}
	personList = append(personList, &p1, &p2, &p3)

	// fmt.Printf("Name: %v\nAge: %v\n", p1.Name, p1.Age)
	
	wg.Add(2)
	go person.UpdatePersonName(&wg, &p1, "Yamada Hanako")
	go person.UpdatePersonAge(&wg, &p1, 22)
	wg.Wait()

	personList = person.RemovePerson(personList, 0)
	
	// fmt.Printf("Name: %v\nAge: %v\n", p1.Name, p1.Age)
	// fmt.Println(personList)
	for _, v := range personList {
		fmt.Println(*v)
	}

}