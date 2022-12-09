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
// 	seconds := time.Now().Unix()
// 	rand.Seed(seconds)
// 	target := rand.Intn(100) + 1
// 	fmt.Println("0부터 100까지의 값을 추출하세요")

//		reader := bufio.NewReader(os.Stdin)
//		for i := 0; i < 10; i++ {
//			fmt.Print("값 입력: ")
//			input, _ := reader.ReadString('\n')
//			input = strings.TrimSpace(input)
//			guess, _ := strconv.Atoi(input)
//			fmt.Println(guess)
//		}
//	}
package main

import "fmt"

type Soldier struct {
	Name  string
	Id    int
	Grade string
}

func NewSoldier(name string, id int, grade string) *Soldier {
	// return &Soldier{name, id, grade}

	ns := new(Soldier)
	ns.Name = name
	ns.Id = id
	ns.Grade = grade
	return ns
}

func main() {
	var s1 = NewSoldier("성윤모", 1234, "이병")
	var s2 = NewSoldier("박민석", 1234, "일병")

	fmt.Printf("%s, %s!\n", s1.Grade, s1.Name)
	fmt.Printf("%s %s의 군번은 %d입니다.\n", s2.Name, s2.Grade, s2.Id)
}
