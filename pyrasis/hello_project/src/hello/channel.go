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
/////////////////////////////////////////////////////////////////////////
/*
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
*/
/////////////////////////////////////////////////////////////////////////
/*
// 채널 버퍼링(비동기)
//다음은 채널의 버퍼가 가득차면 값을 꺼내서 출력합니다.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) // CPU 코어 1개만 사용

	done := make(chan bool, 2) //버퍼 2개인 비동기 채널 생성
	count := 4                 // 반복할 횟수

	go func() {
		for i := 0; i < count; i++ {
			done <- true //채널에 공간이 있으면 true를 보냄, 버퍼가 가득차면 대기
			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done // 버퍼에 값이 없으면 대기, 있으면 값을 꺼냄
		fmt.Println("메인 함수 : ", i)
	}
}
*/
/////////////////////////////////////////////////////////////
/*
// range와 close 사용하기

//range와 close 함수의 특징입니다.
// 이미 닫힌 채널에 값을 보내면 패닉이 발생합니다.
// 채널을 닫으면 range 루프가 종료됩니다.
// 채널이 열려있고, 값이 들어오지 않는다면 range는 실행되지 않고 계속 대기합니다. 만약 다른 곳에서 채널에 값을 보냈다면(채널에 값이 들어오면) 그때부터 range가 계속 반복됩니다.
// 정리하자면 range는 채널에 값이 몇 개나 들어올지 모르기 때문에 값이 들어올 때마다 계속 꺼내기 위해 사용합니다.

//다음은 0부터 4까지 채널에 값을 보내고, 다시 채널에서 값을 꺼내서 출력합니다.

package main

import "fmt"

func main() {
	c := make(chan int) //int형 채널을 생성

	go func() {
		c <- 1
	}()
	a, ok := <-c       // 채널이 닫혔는지 확인
	fmt.Println(a, ok) // 1 true: 채널은 열려 있고 1을 꺼냄

	go func() {
		for i := 0; i < 5; i++ {
			c <- i //채널에 값을 보냄
		}

		close(c) //채널 닫음
	}()

	for i := range c { //range를 사용하여 채널에서 값을 꺼냄
		fmt.Println(i) //꺼낸 값을 출력
	}

	a, ok = <-c        // 채널이 닫혔는지 확인
	fmt.Println(a, ok) // 0 false: 채널이 닫혔음

}
*/
////////////////////////////////////////////////////////////////////////
/*
// 보내기 전용 및 받기 전용 채널 사용하기
package main

import "fmt"

func producer(c chan<- int) { // 보내기 전용 채널
	for i := 0; i < 5; i++ {
		c <- i

	}
	c <- 100

	//fmt.Println(<-c) // 채널에서 값을 꺼내면 컴파일 에러
}

func consumer(c <-chan int) { // 받기 전용 채널
	for v := range c {
		fmt.Println("전달 받은 값 : ", v)
	}
	fmt.Println(<-c) // 채널에 값을 꺼냄

	// c <- 1        // 채널에 값을 보내면 컴파일 에러
}

func main() {

	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
*/
////////////////////////////////////////////////////////////////////////////
/*
// 채널을 리턴값으로 사용
package main

import "fmt"

//                    ↓ 함수의 리턴 값은 int 형 받기 전용 채널
func sum(a, b int) <-chan int {
	out := make(chan int) //채널 생성
	go func() {
		out <- a + b
	}()
	return out
}
func main() {
	c := sum(1, 2) //채널을 리턴값으로 받아서 c에 대입

	fmt.Println(<-c) //3: 채널에서 값을 꺼냄
}
*/
///////////////////////////////////////////////////////////////////////////
/*
// 채널만 사용하여 값 더하기

package main

import "fmt"

//                    ↓ 함수의 리턴 값은 int 형 받기 전용 채널
func num(a, b int) <-chan int {
	c := make(chan int)

	go func() {
		c <- a
		c <- b
		close(c) // 채널을 닫음
	}()

	return c // 채널 변수 자체를 리턴
}

//            ↓ 함수의 매개변수는 int형 받기 전용 채널
func sum(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		r := 0
		for i := range c { // range를 사용하여 채널이 닫힐 때까지 값을 꺼냄 (채널을 꼭 닫아야함)
			r += i
		}
		out <- r
	}()

	return out // 채널 변수 자체를 리턴
}

func main() {
	c := num(1, 2) //1과 2가 들어있는 채널이 리턴됨
	out := sum(c)  //채널 c를 매개변수에 넘겨서 모두 더함, 더한 값이 들어있는 out 채널을 리턴

	fmt.Println(<-out) //3: out 채널에서 값을 꺼냄
}
*/
////////////////////////////////////////////////////////////////////////////
/*
// 셀렉트(select) 사용하기

//Go 언어는 여러 채널을 손쉽게 사용할 수 있도록 select 분기문을 제공합니다.
//'select{ case <-채널: 코드}'

//select 분기문은 switch 분기문과 비슷하지만,
//select 키워드 뒤에 검사할 변수를 따로 지정하지 않으며 각 채널에 값이 들어오면 해당 case가 실행됩니다 (close 함수로 채널을 닫았을 때도 case가 실행됩니다).
//그리고 보통 select를 계속 처리할 수 있도록 for로 반복해줍니다 (반복하지 않으면 한 번만 실행하고 끝냅니다).

//switch 분기문과 마찬가지로 select 분기문도 default 케이스를 지정할 수 있으며
//case에 지정된 채널에 값이 들어오지 않았을 때 즉시 실행됩니다.
//단, default에 적절한 처리를 하지 않으면 CPU 코어를 모두 점유하므로 주의합니다.

// select {
// case <-채널1:
//	 // 채널1에 값이 들어왔을 때 실행할 코드를 작성합니다.
// case <-채널2:
// 	 // 채널2에 값이 들어왔을 때 실행할 코드를 작성합니다.
// default:
// 	 // 모든 case의 채널에 값이 들어오지 않았을 때 실행할 코드를 작성합니다.
// }

//다음은 채널 2개를 생성하고 100밀리초, 500밀리초 간격으로 숫자와 문자열을 보낸 뒤 꺼내서 출력합니다.
package main

import (
	"fmt"
	"time"
)

func main() {
	//채널 생성
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- 10
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1: // 채널 c1에 값이 들어왔다면 값을 꺼내서 i에 대입
				fmt.Println("c1 : ", i)
			case s := <-c2: // 채널 c2에 값이 들어왔다면 값을 꺼내서 s에 대입
				fmt.Println("c2 : ", s)
			case <-time.After(500 * time.Millisecond): // 500 밀리초 후 현재 시간이 담긴 채널이 리턴됨 -> 500 밀리초동안 select안의 case 중 하나라도 호출 되지 않을 시 동작
				fmt.Println("timeout")
			}
			//			//만약 꺼낸 값을 사용하지 않는다면 case <-c1:처럼 변수를 생략해도 됩니다.
			//			select {
			//			case i := <-c1:                // 채널 c1에 값이 들어왔다면 값을 꺼내서 i에 대입
			//				fmt.Println("c1 :", i) // i 값을 출력
			//			case s := <-c2:                // 채널 c2에 값이 들어왔다면 값을 꺼내서 s에 대입
			//				fmt.Println("c2 :", s) // s 값을 출력
			//			}
		}
	}()

	time.Sleep(5 * time.Second)
}
*/
///////////////////////////////////////////////////////////////////////////
/*
// select 분기문은 채널에 값을 보낼 수도 있습니다.
//'case 채널 <- 값: 코드'
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			i := <-c1 //채널 c1에 값을 꺼낸 뒤 i에 대입
			fmt.Println("c1:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world" //채널 c2에 Hello, world를 보냄
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			//여기서는 select에서 매번 채널 c1에 값을 보내지만 채널 c2에 값이 들어오면 c2에서 값을 꺼내서 출력합니다.
			select {
			case c1 <- 10: // 매번 채널 c1에 10을 보냄, 하지만 채널에 값이 들어왔을 때는 값을 받는 case가 실행됩니다.
			case s := <-c2: // c2에 값이 들어왔을 때는 값을 꺼낸 뒤 s에 대입
				fmt.Println("c2:", s) // s값을 출력
			}
		}
	}()

	time.Sleep(5 * time.Second)

}
*/
///////////////////////////////////////////////////////////////////////////

//다음과 같이 채널 c1 한 개로 select에서 값을 보내거나 받을 수도 있습니다.
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)

	go func() {
		for {
			i := <-c1 //채널 c1에 값을 꺼낸 뒤 i에 대입
			fmt.Println("c1:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c1 <- 20 //채널 c2에 20을 보냄
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			//여기서는 매번 채널에 값을 보내지만, select 분기문이 아닌 다른 쪽에서 채널에 값을 보내서 값이 들어왔다면 값을 받는 case가 실행됩니다.
			select {
			case c1 <- 10: // 매번 채널 c1에 10을 보냄, 하지만 채널에 값이 들어왔을 때는 값을 받는 case가 실행됩니다.
			case v := <-c1: // c1에 값이 들어왔을 때는 값을 꺼낸 뒤 v에 대입
				fmt.Println("c1:", v) // v값을 출력
			}
		}
	}()

	time.Sleep(5 * time.Second)

}
