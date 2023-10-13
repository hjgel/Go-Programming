package main

import "fmt"

func paintNeeded(width float64, height float64) float64 { // paintNeeded 함수가 부동 소수점 숫자를 반환하도록 선언
	area := width * height
	return area / 10.0 // 면적 값을 출력하지 않고 반환
}

func main() {
	var amount, total float64 // 현재 계산하고 있는 벽에 필요한 페인트의 양과 페인트의 총량을 저장하는 변수 선언
	amount = paintNeeded(4.2, 3.0) // paintNeeded를 호출한 뒤 반환 값 저장
	fmt.Printf("%0.2f liters neeeded\n", amount) //필요한 페인트 양 출력
	total += amount // 페인트의 양을 페인트의 총량에 더함
	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters neeeded\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)

}
