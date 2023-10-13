package main

import "fmt"

func double(number float64) float64 {
	return number * 2
}
func main() {
	dozen := double(6.0) // 반환 값을 변수에 할당.
	fmt.Println(dozen)
	fmt.Println(double(4.2)) // 반환값을 다른 함수에 전달.
}
