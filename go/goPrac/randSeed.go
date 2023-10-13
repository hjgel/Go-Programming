package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := time.Now().Unix() // 현재 날짜 및 시간을 정숫값으로 가져옴
	rand.Seed(seconds)           // 난수 생성기를 시딩
	target := rand.Intn(100) + 1 // 매번 다른 난수가 생성됨.
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)
}
