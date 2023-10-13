package main

import (
	"bufio"
	"fmt"
	"log" // Fatal function 함수를 사용하기 위함. 에러메시지를 출력하고 프로그램 종료하는 함수
	"os"
	"strconv"
	"strings" // ParseInt
)

func main() {
	fmt.Print("단 입력 : ")
	rd := bufio.NewReader(os.Stdin) // 읽어들이는 함수
	in, err := rd.ReadString('\n')  // 입력받을 변수랑, 에러를 잡을 벼눗 두개가 필요해서 이렇게 선언함

	if err != nil { // 에러가 발생하면
		log.Fatal(err)
	}
	// fmt.Println(in)

	in = strings.TrimSpace(in)
	// dan, err := strconv.ParseInt(in, 10, 32) // in을 10진수로 int32비트로 한다는 개념
	dan, err := strconv.Atoi(in) // ParseInt와 다르게 간단한건 알아서 해주는 Atoi 사용
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 10; i++ {
		// fmt.Println(dan, " * ", i, " = ", (dan * i))
		fmt.Printf("%d * %d = %d\n", dan, i, dan*i)
	}
	// fmt.Println(dan * 2)

	// 다른 언어의 while문 구현
	// i := 1
	// for i < 10 {
	// 	fmt.Print("%d * %d = %d\n", dan, i, dan*i)
	// 	i++
	// }

}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Print("숫자 입력 : ")
// 	rd := bufio.NewReader(os.Stdin) // 읽어들이는 함수
// 	in, _ := rd.ReadString('\n')  // 입력받을 변수랑, 에러를 잡을 벼눗 두개가 필요해서 이렇게 선언함
// 	fmt.Println(in)

// }
