package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" // ParseFloat을 사용하기 위해 strconv 패키지를 가져옴
	"strings" // TrimSpace를 사용하기 위해 strings 패키지를 가져옴
)

func main() {  // 프로그램이 시작될 때 호출함.
	fmt.Print("Enter a grade: ") // 사용자에게 백분율 성적 입력 프롬프트를 출력함
	reader := bufio.NewReader(os.Stdin) // 키보드 입력을 읽기 위한 bufio.NewReader를 생성
	input, err := reader.ReadString('\n') // 사용자가 엔터 키를 누를 때까지 입력한 내용을 읽음
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input) //입력 값에서 줄 바꿈 문자를 제거
	grade, err := strconv.ParseFloat(input, 64) // 입력 문자열을 float64 값으로 변환
	if err != nil {
		log.Fatal(err)
	}

	var status string // status 변수는 함수 스코프에서 사용할 수 있도록 밖에다 선언함
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status) // if문 안에서 선언하는 것 보다 밖에서 해야 오류가 안 남.

}
