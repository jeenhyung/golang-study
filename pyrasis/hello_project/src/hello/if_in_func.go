package main

import "fmt"
import "io/ioutil"

func main() {

	// case1(정석)
	//	var b []byte
	//	var err error

	//	b, err = ioutil.ReadFile("./hello.txt")

	//	if err == nil {
	//		fmt.Printf("%s", b)
	//	}

	// case2
	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		// if 조건문 안에서 변수를 생성할 시, else, if else 문에서는 변수에 접근이 가능하지만,
		// if문 바깥에서 변수에 접근하지 못한다.
		fmt.Printf("%s", b)
	}

	//	fmt.Println(b)   // 변수 b를 사용할 수 없음. 컴파일 에러
	//	fmt.Println(err) // 변수 err을 사용할 수 없음. 컴파일 에러

}
