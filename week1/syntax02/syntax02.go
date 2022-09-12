package main

import (
	"fmt"
	//	"reflect"
	"strings"
)

func main() {
	texts := "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)
	// // 빌드 파일을 만들지 않고 간단하게 빠르게 테스트 하기 위해서는 go run을 써야함 사용 방법으로는
	// // 1. go run ./파일 이름

	// // zero value 초기화하지 않으면 0이 담겨있음
	// // 변수명은 영문자로 시작해야함
	// // 대문자의 경우 다른 패키지에서 접근할 수 있다. 소문자로 시작하는 변수는 동일 패키지에서 접근 가능하다
	// var a int
	// var b float64
	// var c rune
	// var d bool
	// //string type은 빈 값으로 나옴
	// var e string

	// // naming convention
	// // 이것은 관례로 student_id 이런 식으로 쓰지는 않음 뒤에 단어를 대문자로 씀
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

	// fmt.Println(reflect.TypeOf(d))
	// fmt.Println(reflect.TypeOf(e))
}
