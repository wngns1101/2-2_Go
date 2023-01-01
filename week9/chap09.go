// // call by reference
package main

import "fmt"

type Number int

func (n *Number) Double() {
	*n = *n * 2
}

func main() {
	number := Number(4)
	fmt.Println("원래 수:", number)
	number.Double()
	fmt.Println(number)
}
