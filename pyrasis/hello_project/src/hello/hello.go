package main

import (
	"calc"
	"fmt"
	//만약 패키지 디렉터리가 GOPATH/src/hello/calc라면 다음과 같이 사용합니다.
	//즉 기준이 되는 디렉터리는 GOPATH/src입니다.용
	//"hello/calc"
)

func main() {
	fmt.Println("hello, word!")

	fmt.Println(calc.Sum(1, 2)) // calc 패키지의 Sum 함수 사용
}
