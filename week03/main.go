package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	// var a int = 7 // declaration
	// a = 7         // assign value
	// fmt.Println(a)

	a := 7
	fmt.Println(a, reflect.TypeOf(a))

	// b := 8.34 // float64
	var b float32 = 8.34
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(a * int(b))
	fmt.Println(a > int(b))

	var c string // 변수명은 알파벳 대소문자로 시작해야 한다.
	var d int
	var e bool
	var f float64

	koreanZombie := "정찬성"
	fmt.Println(koreanZombie) // camel케이스 사용

	fmt.Println(c, d, e, f)

	fmt.Println('Z', '2', '\n', '김', '인', '히') // rune literals(int32) 90, 50, 10, 44608, 51064
	fmt.Println(reflect.TypeOf("Z"), reflect.TypeOf(2), reflect.TypeOf("Hi"), reflect.TypeOf(4.99), reflect.TypeOf(false))
	fmt.Println(math.Floor(2.17), math.Ceil(2.17), math.Sqrt(16))
	// strings.Title("오픈소스 프로그래밍")
	fmt.Println(strings.Title("open source programming go!"))
	// fmt.Println("OpenSource Programming~", "Go")
}
