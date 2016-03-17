//클로저(Closure)

// Go 언어는 클로저(Closure)를 지원합니다.
// 클로저는 함수 안에서 함수를 선언 및 정의할 수 있고, 바깥쪽 함수에 선언된 변수에도 접근할 수 있는 함수를 말합니다.

package main

import "fmt"

func main() {

	// 익명 함수 사용 기본
	// 익명 함수는 함수를 정의 할 때 이름이 없습니다.
	sum := func(a, b int) int { // <- 익명 함수
		return a + b
	}

	r := sum(1, 2)
	fmt.Println(r) // 3

	// 이름 없는 익명 함수 바깥에 있는 변수 사용
	a, b := 3, 5
	f := func(x int) int {
		return a*x + b // 함수 바깥의 변수 a, b 사용
	}
	y := f(5)
	fmt.Println(y) // 20

	// 그렇다면, 왜 클로저를 사용할까요?
	a, b := 3, 5
	f := calc()
	fmt.Println(f(1)) // 8
	fmt.Println(f(2)) // 11
	fmt.Println(f(3)) // 14
	fmt.Println(f(4)) // 17
	fmt.Println(f(5)) // 20

}

//익명 함수를 리턴하는 calc함수
func calc() func(x int) int {
	a, b := 3, 5 // 지역 변수는 함수가 끝나면 소멸되지만,
	return func(x int) int {
		return a*x + b // *클로저이므로 함수를 호출 할 때마다 변수 a와 b의 값을 사용할 수 있음
	}
}
