package main

import (
	"fmt"
	"reflect"
)

func main() {

	var quantity int
	quantity = 10
	// var l, w float64
	// l, w = 1.2, 2.4
	// var l, w float64 = 1.2, 2.4
	l, w := 1.2, 2.4
	var test bool
	fmt.Println(test)

	fmt.Printf("%.1f %.1f\n", l, w)
	fmt.Println(l*w, "곱셈")

	fmt.Println(reflect.TypeOf(l))

	fmt.Println(quantity)
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1425))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello"))
}
