// package main

// import "fmt"

// func printPointer(myboolPointer *bool) {
// 	fmt.Println(&myboolPointer)
// }
// func main() {
// 	myBool := true
// 	fmt.Println(&myBool)
// }

// package main

// import "fmt"

// func main() {
// 	notes := [3]string{"do", "ra", "mi"}
// 	// for i := 0; i <= 3; i++ {
// 	// for i := 0; i<len(notes); i++{
// 	// for i, note := range notes { 1, do 2, ra 3, mi
// 	// 	fmt.Println(i, note)
// 	// }

// 	// for note := range notes { 출력 1, 2, 3
// 	// 	fmt.Println(note)
// 	// }

// 	for _, note := range notes { //출력 do, ra, mi
// 		fmt.Println(note)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	numbers := [3]float64{71.8, 56.2, 89.5}
// 	sum := 0.0
// 	for _, number := range numbers {
// 		sum += number
// 	}
// 	fmt.Println(sum)
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	f, err := os.Open("data.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scanner := bufio.NewScanner(f)

// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}

// 	err = f.Close()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if scanner.Err() != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFloat(fn string) ([3]float64, error) {
	var numbers [3]float64
	f, err := os.Open(fn)

	if err != nil {
		return numbers, err
	}

	i := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}

	err = f.Close()

	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}
	return numbers, nil
}

func main() {
	prices, err := getFloat("data.txt")
	if err != nil {
		log.Fatal()
	}
	for _, price := range prices {
		fmt.Println(price)
	}
}
