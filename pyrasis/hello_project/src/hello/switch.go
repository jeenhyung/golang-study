package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//case1
	i := 3

	switch i { // 값을 판단할 변수 설정
	case 0: // 각 조건에 일치하는
		fmt.Println(0) // 코드를 실행합니다.
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3: // 3과 변수의 값이 일치하므로
		fmt.Println(3) // 이 부분을 실행하고 이후 실행을 중단
	case 4:
		fmt.Println(4)
	default: // 모든 case에 해당하지 않을 때 실행
		fmt.Println(-1)
	}

	//case2
	s := "world"

	switch s { // 값을 판단할 변수 설정
	case "hello": // 각 조건에 일치하는
		fmt.Println("hello") // 코드를 실행합니다.
	case "world": // 문자열 "world"와 변수의 값이 일치하므로
		fmt.Println("world") // 이 부분을 실행하고 이후 실행을 중단
	default:
		fmt.Println("일치하는 문자열이 없습니다.")
	}

	//case3
	s := "Hello"
	i := 2

	switch i { // 값을 판단할 변수 설정
	case 1:
		fmt.Println(1)
	case 2: // i가 2이고
		if s == "Hello" { // s가 "Hello"이면
			fmt.Println("Hello 2") // Hello 2를 출력하고
			break                  // switch 분기문 실행을 중단
		}

		fmt.Println(2)
	}

	//case4
	i := 3

	switch i { // 값을 판단할 변수 설정
	case 4: // 각 조건에 일치하는
		fmt.Println("4 이상") // 코드를 실행합니다.
		fallthrough
	case 3: // 3과 변수의 값이 일치하므로
		fmt.Println("3 이상") // 이 부분을 실행
		fallthrough         // fallthrough를 사용했으므로 아래 case를 모두 실행
	case 2:
		fmt.Println("2 이상") // 실행
		fallthrough
	case 1:
		fmt.Println("1 이상") // 실행
		fallthrough
	case 0:
		fmt.Println("0 이상") // 실행, 마지막 case에는 fallthrough를 사용할 수 없음
	}

	//case5
	i := 3

	switch i {
	case 2, 4, 6: // i가 2, 4, 6일 때
		fmt.Println("짝수")
	case 1, 3, 5: // i가 1, 3, 5일 때
		fmt.Println("홀수")
	}

	//case6
	rand.Seed(time.Now().UnixNano()) // 현재 시간으로 Seed 값 설정
	switch i := rand.Intn(10); {     // rand.Intn 함수를 실행한 뒤 i에 대입
	case i >= 3 && i < 6: // i가 3보다 크거나 같으면서 6보다 작을 때
		fmt.Println("3 이상, 6 미만") // 코드 실행
	case i == 9: // i가 9일 때
		fmt.Println("9") // 코드 실행
	default: // 모든 case에 해당하지 않을 때
		fmt.Println(i) // 코드 실행
	}

}
