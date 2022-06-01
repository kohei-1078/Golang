package main

import (
	"fmt"
)

// スライス
// 宣言、操作

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	// for
	



	// マップ
	// var m = map[string]int{"A": 100, "B": 200}
	// fmt.Println(m) 

	// m2 := map[string]int{"A": 100, "B": 200}
	// fmt.Println(m2)
	
	// m3 := map[int]string{
	// 	1: "A",
	// 	2: "B",
	// }
	// fmt.Println(m3)

	// m4 := make(map[int]string)
	// fmt.Println(m4)

	// m4[1] = "JAPAN"
	// m4[2] = "USA"
	// m4[4] = "CHINA"

	// fmt.Println(m4)

	// fmt.Println(m["A"])
	// fmt.Println(m4[2])
	// fmt.Println(m4[3])

	// s, ok := m4[1]
	// fmt.Println(s, ok)

	// s2, not := m4[3]
	// fmt.Println(s2, not)

	// delete(m4, 4)
	// fmt.Println(m4)

	// fmt.Println(len(m4))



	// 可変長引数
	// fmt.Println(Sum(1, 2, 3))
	// fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// fmt.Println(Sum())

	// sl := []int{1, 2, 3}
	// fmt.Println(Sum(sl...))

	// for
	// sl := []string{"A", "B", "C"}
	// fmt.Println(sl)

	
	// // for i, v := range sl {
	// 	fmt.Println(i, v)
	// }
// 
	// for i := 0; i < len(sl); i++ {
	// 	fmt.Println(sl[i])
	// }


	// copy
	// sl := []int{100, 200}
	// sl2 := sl
	
	// sl2[0] = 1000
	// fmt.Println(sl)

	// sl := []int{1, 2, 3, 4, 5}
	// sl2 := make([]int, 5, 10)
	// n := copy(sl2, sl)

	// fmt.Println(n, sl2)



	// append make len cap
	// sl := []int{100, 200}
	// fmt.Println(sl)

	// sl = append(sl, 300)
	// fmt.Println(sl)

	// sl = append(sl, 400, 500, 600)
	// fmt.Println(sl)

	// sl2 := make([]int, 5)
	// fmt.Println(sl2)

	// fmt.Println(len(sl2))

	// fmt.Println(cap(sl2))

	// sl3 := make([]int, 5, 10)
	// fmt.Println(sl3)

	// fmt.Println(len(sl3))

	// fmt.Println(cap(sl3))

	// var sl []int
	// fmt.Println(sl)

	// var sl2 []int = []int{100, 200}
	// fmt.Println(sl2)

	// sl3 := []string{"A", "B"}
	// fmt.Println(sl3)

	// sl4 := make([]int, 5)
	// fmt.Println(sl4)

	// sl2[0] = 1000
	// fmt.Println(sl2)

	// sl5 := []int{1, 2, 3, 4, 5}
	// fmt.Println(sl5)

	// fmt.Println(sl5[0])

	// fmt.Println(sl5[2:4])
	
	// fmt.Println(sl5[:2])

	// fmt.Println(sl5[2:])

	// fmt.Println(sl5[:])

	// fmt.Println(sl5[len(sl5)-1])

	// fmt.Println(sl5[1:len(sl5)-1])


}