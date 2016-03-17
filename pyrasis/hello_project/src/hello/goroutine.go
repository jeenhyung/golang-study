// 고루틴(GoRoutine)

/*
//고루틴(Goroutine)은 함수를 동시에 실행시키는 기능입니다.
//다른 언어의 스레드 생성 방법보다 문법이 간단하고,
//스레드보다 운영체제의 리소스를 적게 사용하므로 많은 수의 고루틴을 쉽게 생성할 수 있습니다.
//'go 함수명'

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//(TIP)
//시간표현
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

func hello() {
	fmt.Println("Hellow, world!")
}

func hello2(n int) {
	r := rand.Intn(100)	// 랜덤한 숫자 생성
	time.Sleep(time.Duration(r))	// 랜덤한 시간 동안 대기
	fmt.Println(n)
}

func main() {
	go hello() // 함수를 고루틴으로 실행

	for i := 0; i < 100; i++ { // 100번 반복하여
		go hello2(i) // 고루틴 100개 생성
	}

	fmt.Scanln() // main 함수가 종료되지 않도록 대기
}

*/

/*
// 멀티코어 활용하기
//Go 언어는 CPU 코어를 한 개만 사용하도록 설정되어 있습니다.
//다음은 시스템의 모든 CPU 코어를 사용하는 방법입니다.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정

	fmt.Println(runtime.GOMAXPROCS(0)) //설정 값 출력 (함수에 0을 넣으면 설정 값은 바꾸지 않으며 현재 설정 값만 리턴합니다.)

	s := "Hello, world"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
*/

/*
// 클로저를 고루틴으로 실행하기
//함수 안에서 클로저를 정의한 뒤 고루틴으로 실행할 수 있습니다.
//예제의 출력 결과를 좀 더 확실하게 표현하기 위해 CPU 코어는 하나만 사용하도록 설정합니다.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) // CPU를 하나만 사용

	s := "Hello, world!"

	for i := 0; i < 100; i++ {

		go func(n int) { // 익명 함수를 고루틴으로 실행(클로저)
			fmt.Println(s, n) // s와 매개변수로 받은 n 값 출력
		}(i) // 반복문의 변수는 매개변수로 넘겨줌
	}

	//일반 클로저는 반복문 안에서 순서대로 실행되지만 고루틴으로 실행한 클로저는 반복문이 끝난 뒤에 고루틴이 실행됩니다.

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(s, i) // 반복문의 변수를
			// 클로저에서 바로 출력
		}()
	}

	fmt.Scanln()

	//Go 언어에서 고루틴의 실행 순서를 보장하려면 동기 채널 등을 사용해야 합니다(‘34.1 동기 채널’ 참조).
}
*/
