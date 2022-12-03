// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	resp, err := http.Get("http://www.inhatc.ac.kr")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close() //main 함수 종료 직전에 연결 해지
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(body))
// }

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	responseSize("http://www.inhatc.ac.kr")
// 	responseSize("http://www.daum.net")
// 	responseSize("http://www.nate.com")
// 	responseSize("http://www.naver.com")
// }

// func responseSize(url string) {
// 	fmt.Println(url, "주소를 가져옵니다")
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close() //main 함수 종료 직전에 연결 해지
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(len(body))
// }

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	go responseSize("http://www.inhatc.ac.kr")
// 	go responseSize("http://www.daum.net")
// 	go responseSize("http://www.nate.com")
// 	go responseSize("http://www.naver.com")
// 	time.Sleep(5 * time.Second)
// }

// func responseSize(url string) {
// 	fmt.Println(url, "주소를 가져옵니다")
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close() //main 함수 종료 직전에 연결 해지
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(len(body))
// }

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	pages := make(chan Page)

	urls := []string{
		"http://www.inhatc.ac.kr",
		"http://www.naver.com",
		"http://www.nate.com",
		"http://www.daum.net",
	}
	for _, url := range urls {
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Println("주소는 %s, 길이는 %d\n", page.URL, page.Size)
	}
}

func responseSize(url string, channel chan Page) {
	fmt.Println(url, "주소를 가져옵니다")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() //main 함수 종료 직전에 연결 해지
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}
