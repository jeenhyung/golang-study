//채널

// 채널 사용하기
//채널(channel)은 고루틴끼리 데이터를 주고 받고,
//실행 흐름을 제어하는 기능입니다.
//모든 타입을 채널로 사용할 수 있습니다. 그리고 채널 자체는 값이 아닌 레퍼런스 타입입니다.
//'make(chan 자로형)'

/*
// 고루틴과 채널을 사용하여 두 정수 값을 더해보겠습니다

package main

import "fmt"

//                         ↓ int형 채널을 매개변수로 받음
func sum(a int, b int, c chan int) {
	c <- a + b
	// ↑ 채널에 값을 보냄
}

func main() {

	var c chan int // chan int형 변수 선언
	c = make(chan int)
	//c := make(chan int) // int형 채널 생성

	//채널을 매개변수로 받는 함수는 반드시 go 키워드를 사용하여 고루틴으로 실행해야 합니다.
	go sum(1, 2, c) // sum을 고루틴으로 실행한 뒤 채널을 매개변수로 넘겨줌

	n := <-c       // 채널에서 값을 꺼낸 뒤 n에 대입
	fmt.Println(n) // 3
	//fmt.Println(<-c)

	// <-c는 채널에서 값이 들어올 때까지 대기합니다.
	// 그리고 채널에 값이 들어오면 대기를 끝내고 다음 코드를 실행합니다.
	// 따라서 채널은 값을 주고 받는 동시에 동기화 역할까지 수행합니다.
}

// 채널 <- : 채널에 값을 보냅니다.
// <- 채널 : 채널에 값이 들어올 때까지 기다린 뒤 값이 들어오면 값을 가져옵니다.
//가져온 값은 :=, =를 사용하여 변수에 대입할 수 있습니다.
//값을 가져오면 다음 코드를 실행합니다.
*/

// 동기 채널

//다음은 고루틴과 메인 함수를 번갈아가면서 실행합니다.

package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool) // 동기 채널 생성
	count := 3              // 반복 횟수

	go func() {
		for i := 0; i < count; i++ {
			done <- true // 고루틴에 true를 보냄, 값을 꺼낼 때까지 대기
			fmt.Println("고루틴 : ", i)

		}
	}()

	for i := 0; i < count; i++ {
		<-done                      // 채널에 값이 들어올 때까지 대기, 값을 꺼냄
		time.Sleep(1 * time.Second) // 1초 대기
		fmt.Println("메인 함수 : ", i)  // 반복문의 변수 출력
	}

	//
}
