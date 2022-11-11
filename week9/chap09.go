// call by reference
package main

import "fmt"

type Number int

func (n *Number) Double() {
	*n = *n * 2
}

func main() {
	number := Number(4)
	fmt.Println("원래 수:", number)
A

// // call by value
// package main

// import "fmt"

// type Number int

// func (n Numbner) Double(){
// 	n = n * 2
// }

// func main(){
// 	number := Number(4)
// 	fmt.Println("원래 수:", number)
// 	number.Double()
// 	fmt.Println("두 배가된 수:" number)
// }
