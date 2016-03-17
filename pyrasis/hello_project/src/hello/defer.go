//defer(지연 호출)
//지연 호출은 특정 함수를 현재 함수가 끝나기 직전에 실행하는 기능입니다.
//사용법:
// defer 함수명()
// defer 함수명(매개변수)

package main

import (
	"fmt"
	"os"
)

func main() {

	//case1
	defer world() // 현재 함수(메인함수)가 끝나기 직전에 호출
	hello()
	hello()
	hello()

	//case2
	HelloWorld()

	//case3
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) // 맨 나중에 지연 호출한 함수가 먼저 실행됩니다.
	}

	//case4
	ReadHello()

}

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func HelloWorld() {
	defer func() {
		fmt.Println("world")
	}()

	func() {
		fmt.Println("Hello")
	}()
}

func ReadHello() {
	file, err := os.Open("hello.txt")
	defer file.Close() // 지연 호출한 file.Close()가 맨 마지막에 호출 됨

	if err != nil {
		fmt.Println(err)
		return // file.Close() 호출
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return // file.Close() 호출 (defer에 의해)
	}

	fmt.Println(string(buf))

	// file.Close() 호출 (defer에 의해)

}
