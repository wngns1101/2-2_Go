// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		now := time.Now()
//		year := now.Year()
//		fmt.Println(now)
//		fmt.Println(year)
//	}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	fmt.Print("성적을 입력하세요: ")
// 	rd := bufio.NewReader(os.Stdin)
// 	input, err := rd.ReadString('\n')
// 	if err != nil {
// 		log.Fatal()
// 	}
// 	input = strings.TrimSpace(input)
// 	scores, err := strconv.ParseFloat(input, 64)

// 	if err != nil {
// 		log.Fatal()
// 	}

// 	if scores >= 60 {
// 		fmt.Println("합격")
// 	} else {
// 		fmt.Println("불합격")
// 	}

// }

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println("0부터 100까지의 수를 입력하세요")
	input := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		score, err := input.ReadString('\n')
		if err != nil {
			log.Fatal()
		}
		score = strings.TrimSpace(score)
		guess, err := strconv.Atoi(score)
		if err != nil {
			log.Fatal()
		}
		if guess == target {
			fmt.Println("정답입니다. 정답은", target, "입니다")
			break
		} else if guess > target {
			fmt.Println("더 작게 부르세요")
		} else {
			fmt.Println("더 크게 부르세요")
		}
	}
}
