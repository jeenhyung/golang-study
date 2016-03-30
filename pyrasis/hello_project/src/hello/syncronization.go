// 동기화 객체 사용하기

//Go 언어에서는 채널 이외에도 고루틴의 실행 흐름을 제어하는 동기화 객체를 제공합니다.

/*
대표적인 동기화(synchronization) 객체는 다음과 같습니다.
- Mutex: 뮤텍스입니다. 상호 배제(mutual exclusion)라고도 하며 여러 스레드(고루틴)에서 공유되는 데이터를 보호할 때 주로 사용합니다.
- RWMutex: 읽기/쓰기 뮤텍스입니다. 읽기와 쓰기 동작을 나누어서 잠금(락)을 걸 수 있습니다.
- Cond: 조건 변수(condition variable)입니다. 대기하고 있는 하나의 객체를 깨울 수도 있고 여러 개를 동시에 깨울 수도 있습니다.
- Once: 특정 함수를 딱 한 번만 실행할 때 사용합니다.
- Pool: 멀티 스레드(고루틴)에서 사용할 수 있는 객체 풀입니다. 자주 사용하는 객체를 풀에 보관했다가 다시 사용합니다.
- WaitGroup: 고루틴이 모두 끝날 때까지 기다리는 기능입니다.
- Atomic: 원자적 연산이라고도 하며 더 이상 쪼갤 수 없는 연산이라는 뜻입니다. 멀티 스레드(고루틴), 멀티코어 환경에서 안전하게 값을 연산하는 기능입니다.
*/

//////////////////////////////////////////////////////////////////////////////
// 뮤텍스 사용하기
//뮤텍스는 여러 고루틴이 공유하는 데이터를 보호할 때 사용하며 sync 패키지에서 제공하는 뮤텍스 구조체와 함수는 다음과 같습니다.
//sync.Mutex
//func (m *Mutex) Lock(): 뮤텍스 잠금
//func (m *Mutex) Unlock(): 뮤텍스 잠금 해제

/*
//다음은 고루틴 두 개에서 각각 1,000번씩 슬라이스에 값을 추가합니다.
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든  CPU 사용

	var data = []int{} // int형 슬라이스 생성

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1) // data 슬라이스에 1을 추가

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보(yield)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보(yield)
		}
	}()

	time.Sleep(2 * time.Second) // 2초 대기
	fmt.Println(len(data))      // data 스라이스의 길이 출력(대략 1800~1990 사이의 값)
	//data 슬라이스에 1을 2,000번 추가했으므로 data의 길이가 2000이 되어야 하는데 그렇지가 않습니다.
	//두 고루틴이 경합을 벌이면서 동시에 data에 접근했기 때문에 append 함수가 정확하게 처리되지 않은 상황입니다.
	//이러한 상황을 경쟁 조건(Race condition)이라고 합니다.
}
*/

/*
// 경쟁 조건과 멀티 코어
//CPU의 코어가 여러 개인 컴퓨터에서는 여러 CPU 코어에서 동시에 공유 데이터에 접근할 수 있으므로 경쟁 조건이 발생합니다.
//
//이제 data 슬라이스를 뮤텍스로 보호해보겠습니다.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock() // 뮤텍스 잠금, data 슬라이스 보호 시작

			data = append(data, 1)

			mutex.Unlock() // 뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()

			data = append(data, 1)

			mutex.Unlock()

			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data))
	//여기서는 data 슬라이스를 보호할 것이므로 두 고루틴 모두 data = append(data, 1) 부분 위 아래로 Lock, Unlock 함수를 사용합니다.
	//이제 실행을 해보면 정확히 2000이 출력됩니다.
}
*/
//////////////////////////////////////////////////////////////////////////////////

// 읽기, 쓰기 뮤텍스 사용하기
//
//읽기, 쓰기 뮤텍스는 읽기 동작과 쓰기 동작을 나누어 잠금(락)을 걸 수 있습니다.
/*
 -읽기 락(Read Lock): 읽기 락끼리는 서로를 막지 않습니다. 하지만 읽기 시도 중에 값이 바뀌면 안 되므로 쓰기 락은 막습니다.
 -쓰기 락(Write Lock): 쓰기 시도 중에 다른 곳에서 이전 값을 읽으면 안 되고, 다른 곳에서 값을 바꾸면 안 되므로 읽기, 쓰기 락 모두 막습니다.
*/

/*
//sync 패키지에서 제공하는 읽기, 쓰기 뮤텍스 구조체와 함수는 다음과 같습니다.
sync.RWMutex
func (rw *RWMutex) Lock(), func (rw *RWMutex) Unlock(): 쓰기 뮤텍스 잠금, 잠금 해제
func (rw *RWMutex) RLock(), func (rw *RWMutex) RUnlock(): 읽기 뮤텍스 잠금 및 잠금 해제
*/

/*
//먼저 읽기 쓰기 뮤텍스를 사용하지 않고 고루틴에서 값을 출력해보겠습니다.
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0

	go func() {
		for i := 0; i < 3; i++ {
			data += 1
			fmt.Println("write :", data)      // 쓰기
			time.Sleep(10 * time.Millisecond) // 10밀리초 대기
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("read 1:", data) // 읽기
			time.Sleep(1 * time.Second)  // 1초 대기
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("read 2:", data) // 읽기
			time.Sleep(2 * time.Second)  // 2초 대기
		}
	}()

	time.Sleep(10 * time.Second)

	//실행해보면 매번 불규칙적인 모양으로 출력될 것이고, 특별한 순서를 찾기 어렵습니다.
}
*/

/*
//이제 읽기, 쓰기 동작 실행이 완벽하게 보장되도록 읽기, 쓰기 뮤텍스를 사용해보겠습니다.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	data2 := 10
	var rwMutex = new(sync.RWMutex) // 읽기, 쓰기 뮤텍스 생성

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()                    // 쓰기 뮤텍스 잠금, 쓰기 보호 시작
			data += 1                         // data에 값 쓰기
			fmt.Println("write 1:", data)     // data 값을 출력
			time.Sleep(10 * time.Millisecond) // 10 밀리초 대기
			rwMutex.Unlock()                  // 쓰기 뮤텍스 잠금 해제, 쓰기 보호 종료
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()                   // 쓰기 뮤텍스 잠금, 쓰기 보호 시작
			data2 += 1                       // data2에 값 쓰기
			fmt.Println("write 2:", data2)   // data2 값을 출력
			time.Sleep(9 * time.Millisecond) // 9 밀리초 대기
			rwMutex.Unlock()                 // 쓰기 뮤텍스 잠금 해제, 쓰기 보호 종료
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()              // 읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("read 1:", data) // data 값을 출력(읽기)
			time.Sleep(1 * time.Second)  // 1초 대기
			rwMutex.RUnlock()            // 읽기 뮤텍스 잠금 해제, 읽기 보호 종료
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 2:", data2)
			time.Sleep(2 * time.Second) // 2초 대기
			rwMutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
	//모든 읽기-락 안의 작업이 끝나야 쓰기 작업이 시작된다.
	//쓰기-락 안의 작업이 끝나야 읽기 작업이 시작되며, 쓰기-락이 여러개인 경우 앞의 쓰기 작업이 모두 완료 되어야만 뒤에 쓰기 작업이 시작된다.

	//즉 read 1, read 2 읽기 동작이 모두 끝나야 write 1 쓰기 동작이 시작됩니다.
	//마찬가지로 쓰기 1 동작이 끝나야 읽기 동작이 시작됩니다.
	//읽기 동작끼리는 서로를 막지 않으므로 항상 동시에 실행됩니다.
	//쓰기 1 작업이 모두 완료한 후, 쓰기 2가 동작 합니다.

	//특히 읽기, 쓰기 뮤텍스는 쓰기 동작보다 읽기 동작이 많을 때 유리합니다.
}
*/

// 조건 변수 사용하기(Cond)
//조건 변수는 대기하고 있는 객체 하나만 깨우거나 여러 개를 동시에 깨울 때 사용합니다.

/*
sync 패키지에서 제공하는 조건 변수의 함수는 다음과 같습니다.
 - sync.Cond
 - func NewCond(l Locker) *Cond: 조건 변수 생성
 - func (c *Cond) Wait(): 고루틴 실행을 멈추고 대기
 - func (c *Cond) Signal(): 대기하고 있는 고루틴 하나만 깨움
 - func (c *Cond) Broadcast(): 대기하고 있는 모든 고루틴을 깨움
*/

/*
//먼저 대기하고 있는 고루틴을 하나씩 깨워보겠습니다.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) { // 고루틴 3개 생성
			mutex.Lock() // 뮤텍스 잠금, cond.Wait() 보호 시작

			c <- true // 채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait() // 조건 변수 대기
			fmt.Println("wait end : ", n)

			mutex.Unlock() // 뮤텍스 잠금 해제, cond.Wait() 보호 종료
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
	}

	for i := 0; i < 3; i++ {
		mutex.Lock() // 뮤텍스 잠금, cond.Signal() 보호 시작

		fmt.Println("signal : ", i)
		cond.Signal() // 대기하고 있는 고루틴을 하나씩 깨움

		mutex.Unlock() // 뮤텍스 잠금 해제, cond.Signal() 보고 종료

	}

	fmt.Scanln()
}
*/
//이번에는 대기하고 있는 모든 고루틴을 깨워보겠습니다.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()

			c <- true
			fmt.Println("wait begin : ", n)
			cond.Wait()
			fmt.Println("wait end : ", n)

			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
	}

	mutex.Lock() // 뮤텍스 잠금, cond.Broadcast() 보호 시작

	fmt.Println("boradcast")
	cond.Broadcast() //대기하고 있는 모든 고루틴을 깨움

	mutex.Unlock() // 뮤텍스 잠금, cond.Broadcast() 보호 종료

	fmt.Scanln()
}
