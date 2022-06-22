package person

import (
	"sync"
)

type MyInt int
type MyString string

type Person struct {
	Name MyString
	Age MyInt
}

type PersonList []*Person

func UpdatePersonName(wg *sync.WaitGroup, person *Person, name MyString) {
	person.Name = name
	wg.Done()
}

func UpdatePersonAge(wg *sync.WaitGroup, person *Person, age MyInt) {
	person.Age = age
	wg.Done()
}

func RemovePerson(personList PersonList, i int) PersonList {
	personList[i] = personList[len(personList)-1]
	return personList[:len(personList)-1]
}