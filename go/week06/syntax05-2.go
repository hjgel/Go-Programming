package main

import (
	"fmt"
	"log"
	"os"
)

// 입력(fmt 패키지의 Scanln())된 수의 소수 판정 프로그램 v0.9
func main() {
	var number int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	if err != nil {
		log.Fatal(err)
	}
	for number < 2 {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
		os.Exit(1)
	}

	isPrime := true
	for i := 2; i < number; i++ { // 1과 number일 때 반복문을 돌지 않음
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime { // 비교연산자 제거
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}

}
