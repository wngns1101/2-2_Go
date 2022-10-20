package main

import (
	"fmt"
	"reflect"
)

func main() {
	l, w := 1.2, 2.4
	fmt.Printf("%.1f %.1f\n", l, w)
	fmt.Println(l*w, "곱셈")

	//위에꺼로 해도 되고 아래껄로 해도 됨
	fmt.Println(reflect.TypeOf(l))
	fmt.Printf("%T\n", l)
}
