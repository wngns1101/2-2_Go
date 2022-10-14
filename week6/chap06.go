package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// )

// func getFloat(fn string) ([]float64, error) { // 리턴타입
// 	var numbers []float64 // 선언
// 	f, err := os.Open(fn)

// 	if err != nil {
// 		return nil, err
// 	}

// 	scanner := bufio.NewScanner(f)

// 	for scanner.Scan() {
// 		number, err := strconv.ParseFloat(scanner.Text(), 64) //지역변수
// 		if err != nil {
// 			return nil, err
// 		}
// 		numbers = append(numbers, number) //배열 추가
// 	}

// 	err = f.Close()

// 	if err != nil {
// 		return nil, err
// 	}

// 	if scanner.Err() != nil {
// 		return nil, scanner.Err()
// 	}
// 	return numbers, nil
// }

// func main() {
// 	numbers, err := getFloat("data.txt")
// 	if err != nil {
// 		log.Fatal()
// 	}
// 	sum := 0.0
// 	for _, number := range numbers {
// 		sum = sum + number
// 	}
// 	fmt.Printf("평균: %0.2f\n", sum/(float64(len(numbers))))
// }
