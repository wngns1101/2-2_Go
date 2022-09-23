// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	//var now time.Time = time.Now()
// 	now := time.Now()
// 	year := now.Year()
// 	fmt.Println(year)
// 	fmt.Println(now.Month())
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		//var now time.Time = time.Now()
//		broken := "G# r#cks!"
//		replacer := strings.NewReplacer("#", "o")
//		fixed := replacer.Replace(broken)
//		fmt.Println(fixed)
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
// 	fmt.Print("성적 입력:")
// 	reader := bufio.NewReader(os.Stdin)
// 	// input := reader.ReadString('\n')

// 	// option #1
// 	// input, _ := reader.ReadString('\n')

// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	input = strings.TrimSpace(input)            // 불 필요한 줄바꿈, 캐리지리턴 등의 이스케이프시퀀스를 제거
// 	score, err := strconv.ParseFloat(input, 64) // 입력된 문자열을 64비트 실수형으로 형변환
// 	if err != nil {                             // 2022/09/23 11:23:22 strconv.ParseFloat: parsing "90\n": invalid syntax
// 		log.Fatal(err)
// 	}

// 	var status string
// 	if score >= 60 {
// 		status = "합격"
// 	} else {
// 		status = "불합격"
// 	}
// 	fmt.Println("점수: ", score, " / ", status)
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
	"time"
)

func main() {
	// rand.Seed(99)
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("1에서 100 사이의 수를 추측하세요")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("남은 기회: ", 10-guesses, "회")
		fmt.Println("추측한 값 입력 : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("정답은 추측하신 값보다 더 큽니다.")
		} else if guess > target {
			fmt.Println("정답은 추측하신 값보다 더 작습니다.")
		} else {
			success = true
			fmt.Println("정답입니다! 축하드려요")
			break
		}
	}
	if !success {
		fmt.Println("주어진 기회를 다 쓰셨습니다. 정답은", target, "입니다")
	}
}
