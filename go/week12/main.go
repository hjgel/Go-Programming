package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := make([]int, 5) // 단축연산자 및 슬라이스 리터럴 사용, 변수 선언과 동시에 메모리 할당, 원소들은 제로값으로 초기화.
	s := []int{0, 0, 0, 0, 0} // 변수 선언과 동시에 메모리 할당, 원소 초기화

	s[4] = 100
	s[0] = 7
	s[1] = 91

	for _, value := range s {
		fmt.Println(value)
	}

	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}

	test := [3]string{"인하", "go", "학생"}
	testS := test[0:2]
	fmt.Println(test, len(test))
	fmt.Println(testS, len(testS))
}
