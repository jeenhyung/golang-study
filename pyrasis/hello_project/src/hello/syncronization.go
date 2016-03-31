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
///////////////////////////////////////////////////////////////////////
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
/*
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
*/
//////////////////////////////////////////////////////////////////////

// 함수를 한 번만 실행하기
//Once를 사용하면 함수를 한 번만 실행할 수 있습니다.
/*
sync 패키지에서 제공하는 Once의 구조체와 함수는 다음과 같습니다.
 - sync.Once
 - func (*Once) Do(f func()): 함수를 한 번만 실행
*/
/*
//다음은 고루틴 안에서 Hello, world!를 출력합니다.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello() {
	fmt.Println("Hello, world")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once) // Once 생성

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("gorutine : ", n)
			//Once는 sync.Once를 할당한 뒤에 Do 함수로 사용합니다.
			//Do 함수에는 실행할 함수 이름을 지정하거나, 클로저 형태로 함수를 지정할 수 있습니다.
			once.Do(hello) // 고루틴은 3개지만, hello 함수를 한 번만 실행
		}(i)
	}

	fmt.Scanln()
}
*/
//////////////////////////////////////////////////////////////////////////

// 풀 사용하기
//풀은 객체(메모리)를 사용한 후 보관해두었다가 다시 사용하게 해주는 기능입니다.
//객체를 반복해서 할당하면 메모리 사용량도 늘어나고, 메모리를 해제해야 하는 가비지 컬렉터에게도 부담이 됩니다.
//즉, 풀은 일종의 캐시라고 할 수 있으며 메모리 할당과 해제 횟수를 줄여 성능을 높이고자 할 때 사용합니다. 그리고 풀은 여러 고루틴에서 동시에 사용할 수 있습니다.

/*
sync 패키지에서 제공하는 풀의 구조체와 함수는 다음과 같습니다.
 - sync.Pool
 - func (p *Pool) Get() interface{}: 풀에 보관된 객체를 가져옴
 - func (p *Pool) Put(x interface{}): 풀에 객체를 보관
*/

/*
//풀을 사용하여 정수 10개짜리 슬라이스를 공유해보겠습니다.
//첫 번째 고루틴 그룹에서는 슬라이스에 랜덤한 숫자를 10개를 저장한 뒤 출력하고,
//두 번째 고루틴 그룹에서는 짝수 10개를 저장한 뒤 출력합니다.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct { // Data 구조체 정의
	tag    string // 풀 태그
	buffer []int  // 데이터 저장용 슬라이스
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//풀은 sync.Pool을 할당한 뒤에 Get, Put 함수로 사용합니다.
	//먼저 다음과 같이 sync.Pool을 할당 한 뒤 New 필드에 초기화 함수를 만들어줍니다.
	pool := sync.Pool{ // 풀 할당
		New: func() interface{} { // Get 함수를 사용했을 때 호출될 함수 정의
			data := new(Data)             // 새 메모리 할당
			data.tag = "new"              // 태그 설정
			data.buffer = make([]int, 10) // 슬라이스 공간 할당
			return data                   // 할당한 메모리(객체) 리턴
		},
	}

	// 슬라이스에 랜덤한 숫자를 10개를 저장한 뒤 출력
	for i := 0; i < 10; i++ {
		go func() {
			//풀에서 Get 함수로 객체를 꺼낸 뒤에는 반드시 Type assertion을 해주어야 합니다.
			//여기서는 New 필드의 함수에서 new(Data)로 메모리를 할당했으므로 포인터 형인 (*Data)로 변환합니다.
			data := pool.Get().(*Data) // 풀에서 *Data 타입으로 데이터를 가져옴

			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100) // 슬라이스에 랜덤 값 저장
			}
			fmt.Println(data)
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			//객체를 사용이 끝났으므로 다시 Put 함수를 사용하여 객체를 풀에 보관합니다.
			pool.Put(data) // 풀에 객체를 보관
		}()
	}

	// 짝수 10개를 저장한 뒤 출력합니다.
	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data) // 풀에서 *Data 타입으로 데이터를 가져옴
			n := 0
			for index := range data.buffer {
				data.buffer[index] = n // 슬라이스에 짝수 저장
				n += 2
			}
			fmt.Println(data)
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			pool.Put(data)    // 풀에 객체를 보관
		}()
	}

	//이처럼 풀을 사용하면 메모리를 효율적으로 관리할 수 있습니다.
	//단, 수명 주기가 짧은 객체는 풀에 적합하지 않습니다.

	fmt.Scanln()

}
*/

////////////////////////////////////////////////////////////////////////

// 대기 그룹 사용하기
//대기 그룹은 고루틴이 모두 끝날 때까지 기다릴 때 사용합니다.

/*
sync 패키지에서 제공하는 대기 그룹의 구조체와 함수는 다음과 같습니다.
 - sync.WaitGroup
 - func (wg *WaitGroup) Add(delta int): 대기 그룹에 고루틴 개수 추가
 - func (wg *WaitGroup) Done(): 고루틴이 끝났다는 것을 알려줄 때 사용
 - func (wg *WaitGroup) Wait(): 모든 고루틴이 끝날 때까지 기다림
*/

//Add 함수에 설정한 값과 Done 함수가 호출되는 횟수는 같아야 합니다.
//즉 Add(3)으로 설정했다면 Done 함수는 3번 호출되야 합니다. 이 횟수가 맞지 않으면 패닉이 발생하므로 주의합니다.

/*
//이번에는 대기 그룹을 사용하여 고루틴이 끝날 때까지 기다려보겠습니다
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup) // 대기 그룹 생성

	for i := 0; i < 10; i++ {
		wg.Add(1) // 반복할 때마다 wg.Add 함수로 1씩 추가
		go func(n int) {
			//defer wg.Done() // 고루틴이 끝나기 직전에 wg.Done 함수 호출
			fmt.Println(n)
			wg.Done() // 고루틴이 끝났다는 것을 알려줌
		}(i)
	}

	wg.Wait() // 모든 고루틴이 끝날 떄까지 기다림
	fmt.Println("the end")
}
*/

///////////////////////////////////////////////////////////////////////////

// 원자적 연산 사용하기
//원자적 연산은 더 이상 쪼갤 수 없는 연산이라는 뜻입니다.
//따라서 여러 스레드(고루틴), CPU 코어에서 같은 변수(메모리)를 수정할 때 서로 영향을 받지 않고 안전하게 연산할 수 있습니다.
//보통 원자적 연산은 CPU의 명령어를 직접 사용하여 구현되어 있습니다.

/*
//고루틴을 사용하여 정수형 변수를 2,000번은 더하고, 1,000번은 빼보겠습니다.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			data += 1
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			data -= 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)

	//실행해보면 0 + 2000 - 1000은 1000이 되어야하는데 그렇지가 않습니다(실행할 때마다, 시스템마다 실행 결과는 달라질 수 있습니다).
	//여러 변수에 고루틴이 동시에 접근하면서 정확하게 연산이 되지 않았기 때문입니다.
}
*/

/*
다음은 sync/atomic 패키지에서 제공하는 원자적 연산의 종류입니다.
 - Add 계열: 변수에 값을 더하고 결과를 리턴합니다.
 - CompareAndSwap 계열: 변수 A와 B를 비교하여 같으면 C를 대입합니다. 그리고 A와 B가 같으면 true, 다르면 false를 리턴합니다.
 - Load 계열: 변수에서 값을 가져옵니다.
 - Store 계열: 변수에 값을 저장합니다.
 - Swap 계열: 변수에 새 값을 대입하고, 이전 값을 리턴합니다.
*/

//이번에는 원자적 연산을 사용하여 계산해보겠습니다.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, 1) // 원자적 연산으로 1씩 더함
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, 1) // 원자적 연산으로 1씩 뻄
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)
}
