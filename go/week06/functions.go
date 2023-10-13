// // after (multi return)
// package main

// import "fmt"

// func ProcessScore(kor int, eng int, math int) (int, int) {
// 	totalScore := kor + eng + math
// 	average := totalScore / 3

// 	return totalScore, average
// 	// fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다\n", name, totalScore, average)
// }

// func main() {
// 	var t int
// 	var a int

// 	t, a = ProcessScore(100, 90, 93)
// 	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다\n", "홍길동", t, a)
// 	t, a = ProcessScore(89, 91, 92)
// 	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다\n", "홍길순", t, a)
// }
