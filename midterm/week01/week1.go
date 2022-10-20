// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"math"
// 	"os"
// 	"strings"
// )

// func main() {
// 	// go 빌드 파일 만들기
// 	// 1. go mod init 프로젝트이름/폴더
// 	// 2. go build
// 	// 3. ./프로젝트 이름

// 	// go 빌드 파일 만들지 않고 테스트 하기
// 	// 1. go run ./파일이름

// 	fmt.Println("hello World")

// 	var c rune = '가'

// 	// 문자를 그냥 출력하면 아스키코드값으로 출력됨
// 	fmt.Println(c)

// 	// 이것처럼 printf로 포멧 시켜줘야함
// 	fmt.Printf("%c\n", c)

// 	a := 7
// 	var b float64 = 5.3
// 	a = int(b)
// 	// 타입을 변경해도 자료형은 바뀌지 않음
// 	fmt.Println(a)
// 	fmt.Printf("%T\n", a)

// 	fmt.Println(math.Round(2.71))
// 	fmt.Println(math.Floor(2.71))
// 	fmt.Println(math.Ceil(2.71))

// 	// 문자 바꾸기
// 	texts := "G@ M@ney"
// 	fmt.Println(texts)
// 	r := strings.NewReplacer("@", "o")
// 	newTexts := r.Replace(texts)
// 	fmt.Println(newTexts)

// 	// 입력 받기
// 	fmt.Print("숫자 입력: ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Print(in)
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	fmt.Print("문자입력:")
// 	br := bufio.NewReader(os.Stdin)
// 	str, _ := br.ReadString('\n')

// 	r := strings.NewReplacer("훈", "림")
// 	b := r.Replace(str)
// 	fmt.Println(b)
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// fmt.Print("성적 입력: ")
// reader := bufio.NewReader(os.Stdin)
// input, _ := reader.ReadString('\n')
// input = strings.TrimSpace(input)
// score, _ := strconv.ParseFloat(input, 64)
// if score >= 60 {
// 	fmt.Print("합격")
// } else {
// 	fmt.Print("불합격")
// }
// fmt.Print("성적 입력:")
// reader := bufio.NewReader(os.Stdin)
// str, _ := reader.ReadString('\n')
// str = strings.TrimSpace(str)
// score, _ := strconv.Atoi(str)
// fmt.Print(score)

// seconds := time.Now().Unix()
// rand.Seed(seconds)
// target := rand.Intn(100) + 1

// fmt.Println("1~100까지의 숫자를 추측하시오")
// str := bufio.NewReader(os.Stdin)

// success := false

// for i := 0; i < 10; i++ {
// 	guess, _ := str.ReadString('\n')

// 	guess = strings.TrimSpace(guess)
// 	num, _ := strconv.Atoi(guess)

//		if num == target {
//			success = true
//			fmt.Println("정답입니다")
//		} else if num > target {
//			fmt.Println("큽니다")
//		} else if num < target {
//			fmt.Println("작습니다")
//		}
//	}
//
//	if success == false {
//		fmt.Println("기회를 다 소진하였습니다")
//	}
package main

import "fmt"

func printPointer(myboolPointer *bool) {
	fmt.Println(*myboolPointer)
}
func main() {
	myBool := true
	printPointer(&myBool)
}
