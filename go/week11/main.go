package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5

	// var primes [3]int = [3]int{2, 3, 5} //배열 리터럴로 primes 배열 초기화
	primes := [3]int{2, 3, 5}
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}
	fmt.Println(test[3])

	// for i := 0; i < 3; i++ {
	// fmt.Println((primes[i]))
	// }

	// fmt.Println(primes[3]) // 컴파일 에러

	// i := 0
	// for i < 3 { // 4로하면? panic: runtime error: index out of range [3] with length 3
	// 	fmt.Println((primes[i]))
	// 	i++
	// }

	// for i := 0; i < len(primes); i++ {
	// 	fmt.Println(primes[i])
	// }

	// for idx, prime := range primes {
	// 	fmt.Println(idx, prime)
	// }
	// // enumirate()와 같이 사용할 수 있는 range

	// for idx, prime := range primes { // 선언하고 변수를 사용 안 함.
	// 	fmt.Println(prime)
	// }

	for _, prime := range primes { // 인덱스 사용 안 한다는 뜻
		fmt.Println(prime)
	}

}
