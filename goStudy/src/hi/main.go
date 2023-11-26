package main

// 1. 같은 코드가 있을 때 좀더 가독성을 올리기 위해 쓰는 것이 패키지.
// 디렉터리에 파일을 하나 더 만들어서 import로 불러온다.

import "greeting"

func main() {
	greeting.Hello()
	greeting.Hi()
}
