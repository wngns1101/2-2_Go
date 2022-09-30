package main

import "fmt"

func main() {
	// a := 4
	// pa := &a

	var a int = 4    //int a  = 4;
	var pa *int = &a // int* pa = &a;

	fmt.Println(*pa)
	fmt.Println(pa)
	fmt.Println(&a)
}
