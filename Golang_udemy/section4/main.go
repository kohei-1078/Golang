package main 

import (
	"fmt"
	"strconv"
)

// 定数

const Pi = 3.14

const (
	URL = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

const (
	c0 = iota
	c1
	c2
)

func main() {
	/*
	// 型
	// 整数型

	var i int = 100
	var i2 int64 = 200

	fmt.Println(i + 50)

	// fmt.Println(i + i2)

	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))

	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	// 論理値型
	var t, f bool = true, false
	fmt.Println(t, f)

	// 文字列型
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test
		test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))

	// byte型
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))
	
	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))

	// 配列型
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	arr2[2] = "C"
	fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	fmt.Println(len(arr1))

	// interface & nil
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)

	x = "A"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// x = 2
	// fmt.Println(x + 3)

	// 型変換
	var i11 int = 1
	fl6411 := float64(i11)
	fmt.Println(fl6411)
	fmt.Printf("i11 = %T\n", i11)
	fmt.Printf("fl6411 = %T\n", fl6411)

	i12 := int(fl6411)
	fmt.Printf("i12 = %T\n", i12)
	
	fl13 := 10.5
	i13 := int(fl13)
	fmt.Printf("i13 = %T\n", i13)
	fmt.Println(i13)
	

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)
*/

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)

	// 定数
	fmt.Println(Pi)
	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
	
	// 算術演算子
	fmt.Println(1 + 2)
	fmt.Println("ABC" + "DE")

	fmt.Println(5 - 1)

	fmt.Println(5 * 4)

	fmt.Println(60 / 3)

	fmt.Println(9 % 4)

	n := 0
	n += 4 // n=n+4
	fmt.Println(n)
	n++ // n=n+1
	fmt.Println(n)
	n-- // n=n-1
	fmt.Println(n)

	s := "ABC"
	s += "CEF"
	fmt.Println(s) //s = "ABC" + "DEF"

	fmt.Println(1 == 1)

	fmt.Println(1 == 2)

	fmt.Println(4 <= 8)
	
	fmt.Println(1 >= 8)

	fmt.Println(1 < 8)
	fmt.Println(3 > 1)

	fmt.Println(true == false)
	fmt.Println(true != false)

	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)

	fmt.Println(true || false == true)
	fmt.Println(false || false == true)

	fmt.Println(!true)
	fmt.Println(!false)





	

}