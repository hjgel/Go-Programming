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

func prime(number int) {
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

func primeRange(a int, b int) {
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
}

// 소수 판정 및 구간 소수 판정 프로그램 v1.7
func main() {
	var menu int

	for true {
		fmt.Print("1) 소수판정 2) 구간 소수판정 0) 종료 : ")
		_, err := fmt.Scanln(&menu)

		if err != nil {
			log.Fatal(err)
		}

		switch menu {
		case 1:
			var in int
			fmt.Print("정수입력 : ")
			_, err := fmt.Scanln(&in)

			if err != nil {
				log.Fatal(err)
			}
			prime(in)
		case 2:
			var n1, n2 int
			fmt.Print("정수 2개 입력 : ")
			_, err := fmt.Scanln(&n1, &n2)

			if err != nil {
				log.Fatal(err)
			}
			primeRange(n1, n2)
		default:
			fmt.Print("프로그램을 종료합니다")
			os.Exit(1)
		}

	}

	// if p { // 비교연산자 제거
	// 	fmt.Println(number, "는(은) 소수입니다.")
	// } else {
	// 	fmt.Println(number, "는(은) 소수가 아닙니다~")
	// }

}
