package main

import (
	"fmt"
)

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

// type error interface {
// 	Error() string
// }

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

type Point struct {
	A int
	B string
}

// func (p *Point) String() string {
// 	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
// }

func main() {
	p := &Point{100, "ABC"}
	fmt.Printf("%T, %v\n", p, p)
	fmt.Println(p)

	// err := RaiseError()
	// fmt.Println(err.Error())

	// e, ok := err.(*MyError)
	// if ok {
	// 	fmt.Println(e.ErrCode)
	// }


	// vs := []Stringfy{
	// 	&Person{Name: "Taro", Age: 21},
	// 	&Car{Number: "123-456", Model: "AB-1234"},
	// }

	// for _, v := range vs {
	// 	fmt.Println(v.ToString())
	// }

}