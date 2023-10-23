package main

import "fmt"

func double(n *int) {
	*n = *n * 2
}

func main() {
	var a int = 5
	double(&a) // pass by pointer
	fmt.Printf("%d\n", a)
}

// package main

// import "fmt"

// func main() {
// 	a := 10 // var a int = 10
// 	var pa *int
// 	pa = &a

// 	fmt.Println(a, *pa)
// 	fmt.Println(&a, pa)
// 	fmt.Printf("%T %T %T %T\n", a, *pa, &a, pa)
// 	fmt.Println(&pa)
// 	// pa는 a의 주소값을 가지고. pa
// 	// pa의 주소는 다른 값을 가짐. &pa
// 	// pa가 가르키는 곳은 a의 값을 가짐. *pa
// }
