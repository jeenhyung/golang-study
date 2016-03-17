// panic
// 프로그램이 잘못되어 에러가 발생한 뒤 종료되는 상황을 패닉이라 합니다.
// 패닉은 panic() 함수를 이용해 임의로 만들 수 있습니다.
// recover 함수를 사용하면 패닉이 발생했을 때 프로그램이 바로 종료되지 않고 예외 처리를 할 수 있으며 다른 언어의 try catch 구문과 비슷하게 동작합니다.

package main

import "fmt"

////case1
//func main() {
//	panic("Error !!")
//	fmt.Println("Hello, world")	//실행되지 않음

//}

////case2
//func main() {
//	f()

//	fmt.Println("hello, world!")

//}

//func f() {
//	defer func() { // recover 함수는 지연 호출로 사용해야 함
//		s := recover() // 패닉이 발생해도 프로그램을 종료하지 않음, panic 함수에서 설정한 에러 메시지를 받아옴
//		fmt.Println(s)
//	}()

//	panic("Error !!!") // panic 함수로 에러 메시지 설정, 패닉 발생
//}

//case3
func main() {
	f()
	fmt.Println("Hello, world!") // 런타임 에러가 발생했지만 recover 함수로 복구되었기 때문에 이 부분은 정상적으로 실행됨
}

func f() {
	defer func() {
		s := recover() // recover 함수로 런타임 에러(패닉) 상황을 복구
		fmt.Println(s)
	}()

	a := [...]int{1, 2} // 정수가 2개 저장된 배열

	for i := 0; i < 5; i++ { //배열 크기를 벗어난 접근
		fmt.Println(a[i])
	}
}
