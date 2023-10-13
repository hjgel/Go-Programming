package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var c rune = '가' // rune은 작은 따옴표?
	// var a = 7
	a := 7.3
	a = 5 // 값은 바뀔 수 있어도, 타입은 바뀌지 않음

	// b := 7
	//b = 5.3 // int에서 float으로 바꾸는 건 안 됨. (형변환으로 하면 됨.)
	b := 7
	var d float64 = 5.3
	b = int(d) // 이런식으로 형변환 가능

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	e := false

	fmt.Println(e)
	fmt.Printf("%T\n", e)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(c)        // 유니코드(UTF-8) 값 출력
	fmt.Printf("%c\n", c) // 한 글자 출력
	fmt.Printf("%T\n", c) // rune은 결국 int32의 별명이다.

	fmt.Println(math.Ceil(2.71))  // 올림 함수
	fmt.Println(math.Floor(2.71)) // 내림 함수
	fmt.Println(math.Round(2.71)) // 반올림 함수
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github java"))
}
