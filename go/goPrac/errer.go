package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 { // 너비 값이 유효하지 않은 경우 0과 에러를 반환함
		return 0, fmt.Errorf("a width of %0.2f is invaild", width)
	}
	if height < 0 { // 높이 값이 유효하지 않은 경우 0과 에러 반환
		return 0, fmt.Errorf("a height of %0.2f is invaild", height)
	}
	area := width * height
	return area / 10.0, nil // 아무 이상이 없으면 없다고 알려주기 위해 nil 반환
}

func main() {
	amount, err := paintNeeded(4.2, -3.0) // err에 height 보낸거임.
	fmt.Println(err)
	fmt.Printf("%0.2f liter needed\n", amount)
}
