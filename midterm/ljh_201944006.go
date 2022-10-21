package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func factorial(x int) int {
	sum := 0
	for i := x; i > 0; i-- {
		sum += i
	}
	return sum
}

func cal(x int) bool {
	find := true
	for i := 2; i < x; i++ {
		if x%i == 0 {
			find = false
		}
	}
	if x == 0 {
		find = false
	}
	return find
}

func main() {
	var numbers []string
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	numbers = strings.Split(input, " ")
	fmt.Println(numbers)
	for _, number := range numbers {
		i, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		if i < 0 {
			continue
		} else {
			find := cal(i)
			num := factorial(i)
			fmt.Println(i, "펙토리얼 수: ", num, "/ 소수 여부 : ", find)
		}
	}
}
