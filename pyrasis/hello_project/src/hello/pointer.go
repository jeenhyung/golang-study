//포인터

package main

import "fmt"

/*
//case1
func main() {

	var numPtr *int     // 포인터형 변수를 선언하면 nil로 초기화됨
	fmt.Println(numPtr) // nil

	// 빈 포인터형 변수는 바로 사용할 수 없으므로, new 함수로 메모리를 할당해야 합니다.
	// Go 언어는 메모리를 관리해주는 가비지 컬렉션을 지원하므로 메모리를 할당한 뒤 해제하지 않아도 됩니다.
	var numPtr2 *int = new(int)
	fmt.Println(numPtrs)

	// 포인터형 변수에 값을 대입하거나, 가져오려면 역참조(dereference)를 사용합니다.
	var numPtr3 *int = new(int) // new 함수로 공간 할당
	*numPtr3 = 1                // 역참조로 포인터형 변수에 값을 대입
	fmt.Println(*numPtr3)       // 1: 포인터형 변수에서 값을 가져오기

	//일반 변수에 참조(레퍼런스)를 사용하면 포인터형 변수에 대입할 수 있습니다.
	var num int = 1
	var numPtr4 *int = &num // 참조로 num 변수의 메모리 주소를 구하여 numPtr4 포인터 변수에 대입
	fmt.Println(numPtr4)    // numPtr4 포인터 변수에 저장된 메모리 주소
	fmt.Println(&num)       // 참조로 num 변수의 메모리 주소를 구함

	// Go 언어에서는 메모리 주소를 직접 대입하거나 포인터 연산을 허용하지 않습니다. 따라서 다음과 같이 메모리 주소를 직접 조작할 수는 없습니다.
	var numPtr5 *int = new(int)
	numPtr5++              // 컴파일 에러. 포인터 연산은 허용하지 않음
	numPtr5 = 0xc0820062d0 // 컴파일 에러. 메모리 주소를 직접 대입할 수 없음
	fmt.Println(numPtr5)
}
*/

//case2
func hello(n int) {
	n = 2 // 매개변수 n에 2를 대입
}

func world(n *int) {
	*n = 2 // 포인터 변수 n을 역참조하여 메모리에 2를 대입
}

func main() {
	var n int = 1

	//일반 함수
	hello(n)       // 1이 들어있는 변수 n을 hello 함수에 넘김
	fmt.Println(n) // 1: hello 함수안의 n에 2를 대입했지만 바깥의 n은 영향이 없음

	//포인터 매개변수를 받는 함수
	world(&n)      // 1이 들어있는 변수 n의 메모리 주소를 world 함수에 넘김
	fmt.Println(n) // 2: world 함수에서 n의 메모리 공간에 2를 대입했으므로 바깥에 있는 n의 값이 바뀌었음

	//내부 함수(클로저)
	func() {
		n = 3
	}()
	fmt.Println(n) // 3: n에 영향이 있음

}
