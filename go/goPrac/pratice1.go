package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dolloars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate // price와 taxRate의 type이 다르기 때문
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax // 이것도 type 문제
	fmt.Println("Total cost is", total, "dollars")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available")
	fmt.Println("Within budget?", total <= float64(availableFunds)) // 이것도 형변환 문제

}
