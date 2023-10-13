package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	prime := true

	if n < 2 {
		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다~", n)
	}

	for i := 2; i < n; i++ { // 1과 number일 때 반복문을 돌지 않음
		if n%i == 0 {
			prime = false
			break
		}
	}
	return prime, nil // 소수 판정 값, 정상 데이터
}

// 소수 판정 프로그램 v1.0 : 함수 적용
func main() {
	var number int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	if err != nil {
		log.Fatal(err)
	}

	p, err := isPrime(number)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if p { // 비교연산자 제거
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}

}
