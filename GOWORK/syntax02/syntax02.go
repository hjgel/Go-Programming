package main

import (
	"fmt"
	"strings"
)

func main() {
	texts := "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)

	fmt.Println(newTexts)

	// 변수명은 영문자로 시작해야 된다.
	// 영문 대문자의 경우 다른 패키지에서 접근할 수 있다. 소문자로 시작하는 변수는 동일 패키지에서만 접근 가능하다.
	// zero value
	// var e string // string은 비어있는 문자열 출력
	// var d bool
	// var c rune
	// var b float64
	// var a int

	// // naming convention
	// var studentId string
	// var i int32
	// // a := 7

	// fmt.Println(studentId)
	// fmt.Println(i)

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// fmt.Printf("%T\n", c)
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", d)
	// fmt.Printf("%T\n", e)

	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(reflect.TypeOf(a))

}
