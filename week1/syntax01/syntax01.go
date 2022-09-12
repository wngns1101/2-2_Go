package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// go는 빌드 파일이 있어야지 실행 가능함 절차로는
	// 1. go mod init 프로젝트명/폴더
	// 2. go build
	// 3. ./파일 이름

	// 변수 선언
	// 변수 선언 변수 이름 변수 타입
	var c rune = '가'

	//다양한 변수 선언 방법
	// var a int16 = 7
	// var a = 7
	// a := 7
	//값은 바뀔 수 있어도 타입은 바뀌지 않음, float에서 int로는 가능하나 int에서 float로 변경은 안 됨
	a := 7
	var b float64 = 5.3
	a = int(b) //casting 연산자를 사용해야함
	d := false

	//유니코드 값 출력
	fmt.Println(c)
	fmt.Println(a)

	//문자 출력
	fmt.Printf("%c\n", c)

	//타입 출력  utf-8은 1~3 byte rune은 int32의 별명(4바이트 사용)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", d)

	fmt.Println(math.Round(2.71))
	fmt.Println(math.Floor(2.71))
	fmt.Println(math.Ceil(2.71))

	fmt.Println("Hello Go~")

	fmt.Println(strings.Title("go git github java"))
}
