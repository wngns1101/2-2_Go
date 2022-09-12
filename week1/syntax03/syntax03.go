package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("숫자 입력 : ")
	// os.Stdin은 표준 입력으로 키보드를 의미
	rd := bufio.NewReader(os.Stdin)
	// readString하면 값을 받을 변수와 에러를 받을 변수가 필요한데 go는 사용을 하지 않으면 저장할 때 사라지기 때문에 _로 대체함
	in, _ := rd.ReadString('\n')
	fmt.Println(in)
}
