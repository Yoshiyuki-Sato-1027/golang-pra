package main

import (
	"fmt"
	// "time"
)
// JSでいうと、、、(齟齬あるかも)
// fmt.Println→console.log
// int→number

func Plus(x,y int)int {
return x + y
}

// (int, int)は返り値の型
func Div(x, y int) (int, int) {
	q := x/y
	r := x % y
	return q,r
}

func Double(price int) (result int) {
	result = price *2
	return result
}

// 返り値なし
func Noreturn() {
	fmt.Println("No Return")
	return 
}

func main() {
// 	fmt.Println("Hello World!")
// 	fmt.Println(time.Now())
// 	var i int = 100
// 	fmt.Println(i)
// 	var s string = "100"
// 	fmt.Println(s)
// 	var t,f bool = true,false
// 	fmt.Println(t,f)
// 	var (
// 		i2 int = 1009
// 		s2 string = "golang"
// 	)
// 	fmt.Println(i2,s2)

// 	var i3 int
// 	var s3 string
// 	fmt.Println(i3,s3)

// 	// 暗黙的な定義
// 	i4 := 400
// 	fmt.Println(i4)

// 	i4 = 450
// 	fmt.Println(i4)

// 	var arr1 [3]int
// 	fmt.Println(arr1)

// 	var arr2 [3]string = [3]string{"A","B"}
// 	fmt.Println(arr2)

// 	arr3 := [3]int{1,2,3}
// 	fmt.Println(arr3)
// // ...は要素数をカウントしてくれる
// 	arr4 := [...]string{"C","D"}
// 	fmt.Println(arr4)

// 	fmt.Println(arr1[0])
// 	fmt.Println(len(arr1))

	i := Plus(1,2)
	fmt.Println(i)
	// _は引数なし
	i2, _ := Div(9,3)
	fmt.Println(i2)
	i4 := Double(1000)
	fmt.Println(i4)
	Noreturn()
}