package main

import (
	"fmt"
)

// shadowing
func main() {
	// 자료 타입을 변수명으로 사용
	var test1 float64 = 9.1
	var test2 float64 = 7.9
	var univ string = "inha"

	var f1 string = "functions"
	var f2 = append([]string{}, "함수")

	fmt.Println(test1 + test2)
	fmt.Println(univ)
	fmt.Println(f1)
	fmt.Println(f2)

	// 패키지를 변수명으로 사용

	// fmt.Println() // 변수명과 겹치면 안 됨.

	// 함수를 변수명으로 사용
	// var append string = "functions"
	// var f = append([]string{}, "함수")
}
