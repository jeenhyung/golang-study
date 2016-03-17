package main

import "fmt"

func main() {

	hello()

	r := sum(1, 2)
	fmt.Println(r)

	rr := sub(1, 2)
	fmt.Println(rr)

	//리턴값 여러개
	_, sub := SumAndSub(6, 2)
	fmt.Println(sub)

	//가변인자

	rrr := sum2(1, 2, 3, 4, 5)
	fmt.Println(rrr)

	n := []int{1, 2, 3, 4, 5}
	rrrr := sum2(n...) // ...을 사용하여 가변인자에 슬라이스를 바로 넘겨줌
	fmt.Println(rrrr)

	//함수를 변수에 저장하기

	var hello func(a int, b int) int = sum // 함수를 저장하는 변수를 선언하고 함수 대입
	world := sum                           // 변수 선언과 동시에 함수를 바로 대입
	fmt.Println(hello(1, 2))               // 3: hello 변수에 저장된 sum 함수 호출
	fmt.Println(world(1, 2))               // 3: world 변수에 저장된 sum 함수 출호

	f := []func(int, int) int{sum, sub} // 함수를 저장할 수 있는 슬라이스를 생성한 뒤 함수로 초기화
	fmt.Println(f[0](1, 2))             // 3: 배열의 첫 번째 요소로 함수 호출
	fmt.Println(f[1](1, 2))             // -1: 배열의 두 번째 요소로 함수 호출

	ff := map[string]func(int, int) int{
		"sum": sum,
		"sub": sub,
	}
	fmt.Println(ff["sum"](1, 2)) // 3: 맵에 sum 키를 지정하여 함수 호출
	fmt.Println(ff["sub"](1, 2)) // -1: 맵에 diff 키를 지정하여 함수 호출

	//익명 함수 사용하기

	// ( 이런 익명 함수는 코드양을 줄일 수 있으며,
	// 클로저, 지연 호출(defer), 고루틴(go)에서 주로 사용 )

	func() { // 함수에 이름이 없음
		fmt.Println("Hello, world!")
	}()

	func(s string) { // 익명 함수를 정의한 뒤
		fmt.Println(s)
	}("Hello, world") // 바로 호출

	r5 := func(a int, b int) int { // 익명 함수를 정의한 뒤
		return a + b
	}(1, 2) // 바로 호출하여 리턴값을 변수 r에 저장

	fmt.Println(r5)

}

func hello() {
	fmt.Println장("hello, world")
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) (r int) {
	r := a - b
	return
}

func SumAndSub(a int, b int) (int, int) {
	return a + b, a - b
}

func SumAndSub(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func sum2(n ...int) int { // int형 가변인자를 받는 함수 정의
	totla := 0
	for _, value := range n { // range로 가변인자의 모든 값을 꺼냄 (n은 슬라이스 타입)
		totla += value // 꺼낸 값을 모두 더함
	}
	return total
}
