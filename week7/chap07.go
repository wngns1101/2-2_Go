package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getStrings(fn string) ([]string, error) { // 리턴타입
	var lines []string
	f, err := os.Open(fn)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = f.Close()

	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func main() {
	lines, err := getStrings("votes.txt")
	if err != nil {
		log.Fatal()
	}

	fmt.Printf("투표용지: %s\n", lines)
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("%s:%d\n", name, count)
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func getStrings(fn string) ([]string, error) { // 리턴타입
// 	var lines []string
// 	f, err := os.Open(fn)

// 	if err != nil {
// 		return nil, err
// 	}

// 	scanner := bufio.NewScanner(f)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		lines = append(lines, line)
// 	}

// 	err = f.Close()

// 	if err != nil {
// 		return nil, err
// 	}

// 	if scanner.Err() != nil {
// 		return nil, scanner.Err()
// 	}
// 	return lines, nil
// }

// func main() {
// 	lines, err := getStrings("votes.txt")
// 	if err != nil {
// 		log.Fatal()
// 	}

// 	fmt.Printf("투표용지: %s\n", lines)
// 	var names []string
// 	var counts []int

// 	for _, line := range lines {
// 		matched := false
// 		for i, name := range names {
// 			if name == line { // 이름이 있으면
// 				counts[i]++
// 				matched = true
// 			}
// 		}
// 		if matched == false { //첫번째 이름이 나올 때만 동작
// 			names = append(names, line)
// 			counts = append(counts, 1)
// 		}
// 	}
// 	for i, name := range names {
// 		fmt.Printf("%s:%d\n", name, counts[i])
// 	}
// }
