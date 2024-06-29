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

func ReturnFunc() func() {
	return func() {
		fmt.Println("function")
	}
}

func CallFunc(f func()) {
f()
}

func main() {

	CallFunc(func() {
		fmt.Println("Func")
	})
}