package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(time.Now())
	var i int = 100
	fmt.Println(i)
	var s string = "100"
	fmt.Println(s)
	var t,f bool = true,false
	fmt.Println(t,f)
	var (
		i2 int = 1009
		s2 string = "golang"
	)
	fmt.Println(i2,s2)

	var i3 int
	var s3 string
	fmt.Println(i3,s3)

	// 暗黙的な定義
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)
}