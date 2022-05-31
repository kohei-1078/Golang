package main

import (
	"fmt"
)

// 関数

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
}

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func CallFunction(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func intergers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// i := Plus(1, 2)
	// fmt.Println(i)
	
	// i2, _ := Div(9, 4)
	// fmt.Println(i2)

	// i4 := Double(1000)
	// fmt.Println(i4)

	// Noreturn()

	// 無名関数

	// f := func(x, y int) int {
	// 	return x + y
	// }
	// i := f(1,2)
	// fmt.Println(i)

	// i2 := func(x, y int) int {
	// 	return x + y
	// }(1, 2)

	// fmt.Println(i2)

	// 関数を返す
	// f := ReturnFunc()
	// f()

	// 関数を引数に取る関数
	// CallFunction(func() {
	// 	fmt.Println("I'm a function")
	// })

// クロージャーの実装
	// f := Later()
	// fmt.Println(f("Hello"))
	// fmt.Println(f("My"))
	// fmt.Println(f("name"))
	// fmt.Println(f("is"))
	// fmt.Println(f("Golang"))

	// ジェネレーターの実装
	ints := intergers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints := intergers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())



}