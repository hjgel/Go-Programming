package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	//prime := true

	if n < 2 {
		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다~", n)
	}

	for i := 2; i < n; i++ { // 1과 number일 때 반복문을 돌지 않음
		if n%i == 0 {
			return false, nil
			// prime = false
			// break
		}
	}

	return true, nil // 소수 판정 값, 정상 데이터
}

// 구간 소수 판정 프로그램 v1.2 isPrime 함수 안의 변수를 하나 줄이고 return 구문 추가, break 제거
func main() {
	var a, b int
	// 2 20
	// 2 3 5 7 11 13 ... 19

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&a, &b)

	if err != nil {
		log.Fatal(err)
	}

	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if p {
			fmt.Print(i, " ")
		}
	}

	// if p { // 비교연산자 제거
	// 	fmt.Println(number, "는(은) 소수입니다.")
	// } else {
	// 	fmt.Println(number, "는(은) 소수가 아닙니다~")
	// }

}
