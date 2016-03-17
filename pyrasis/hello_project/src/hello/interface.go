//인터페이스

/*
//인터페이스는 메서드 집합입니다. 단 인터페이스는 메서드 자체를 구현하지는 않습니다.
//'type 인터페이스명 interface { }'

package main

import "fmt"

type hello interface { //인터페이스 정의

}

func main() {

	var h hello    //인터페이스 선언
	fmt.Println(h) // <nil>: 빈 인터페이스이므로 nil이 출력됨

}
*/

/*
//다음은 int 자료형에 메서드를 연결하고, 인터페이스로 해당 메서드를 호출합니다.

package main

import "fmt"

type MyInt int

func (i MyInt) Print() {
	fmt.Println(i)
}

type Printer interface {
	Print()
}

func main() {

	var i MyInt = 5
	var p Printer	//인터페이스 선언

	p = i	//i를 인터페이스 p에 대입
	p.Print()	// 5: 인터페이스를 통하여 MyInt의 Print 메서드 호출

}
*/

/*
//다음은 int 자료형과 사각형 구조체의 내용을 출력하고,
//int 자료형과 사각형 구조체의 인스턴스를 담을 수 있는 인터페이스를 정의한 예제입니다.

package main

import "fmt"

type MyInt int // int 형을 MyInt로 정의

type Rectangle struct { // 사각형 구조체 정의
	width  int
	height int
}

func (i MyInt) Print() { // MyInt에 Print 메서드를 연결
	fmt.Println(i)
}
func (rect Rectangle) Print() { // Rectangle에 Print 메서드를 연결
	fmt.Println(rect.width, rect.height)
}

type Printer interface { // Print 메서드를 가지는 인터페이스 정의
	Print()
}

func main() {
	var i MyInt = 5
	rect := Rectangle{10, 20}

	var p Printer // 인터페이스 선언

	p = i
	p.Print() // 5: 인터페이스 p를 통하여 MyInt의 Print 메서드 호출

	p = rect
	p.Print() // 10 20: 인터페이스 p를 통하여 Rectangle의 Print 메서드 호출

	//즉 인터페이스는 자료형이든 구조체든 타입에 상관없이 '메서드 집합만 같으면' 동일한 타입으로 봅니다.
	//따라서, 똑같은 타입(매개변수, 이름, 리턴값)의 함수를 가지고 있지 않으면 인터페이스에 대입할 수 없습니다.

	//인터페이스를 선언하면서 초기화하려면 다음과 같이 :=를 사용하면 됩니다.
	//인터페이스에는 ( ) (괄호)를 사용하여 변수나 인스턴스를 넣어줍니다.
	p1 := Printer(i) // 인터페이스를 선언하면서 i로 초기화
	p1.Print()       // 5

	p2 := Printer(rect) // 인터페이스를 선언하면서 rect로 초기화
	p2.Print()       // 10 20

	//다음과 같이 배열(슬라이스) 형태로도 인터페이스를 초기화 할 수 있습니다.
	pArr := []Printer{i, rect}	//슬라이스 형태로 인터페이스 초기화

	for index, _ := range pArr {
		pArr[index].Print()	// 슬라이스를 순회하면서 Print 메서드 호출
	}

	for -, value := range pArr {
		value.Print()	// 슬라이스를 순회하면서 Print 메서드 호출
	}

}
*/

/*
// 덕 타이핑

//이렇게 각 값이나 인스턴스의 실제 타입은 상관하지 않고 구현된 메서드로만 타입을 판단하는 방식을 덕 타이핑(Duck typing)이라 합니다.
//이 용어는 다음과 같은 덕 테스트(오리 테스트)에서 유래되었습니다.
//“만약 어떤 새가 오리처럼 걷고, 헤엄치고, 꽥꽥거리는 소리를 낸다면 나는 그 새를 오리라 부르겠다.”

package main

import "fmt"

type Duck struct { // 오리(Duck) 구조체 정의
}

func (d Duck) quack() { // 오리의 quack 메서드 정의
	fmt.Println("꽥~!")
}
func (d Duck) feather() { // 오리의 feathers 메서드 정의
	fmt.Println("오리는 흰색 회색 털을 가지고 있습니다.")
}

type Person struct { // 사람(Person) 구조체 정의
}

func (p Person) quack() { // 사람의 quack 메서드 정의
	fmt.Println("사람이 오리 흉내를 냅니다. 꽥~!") // 사람이 오리 소리를 흉내냄
}
func (p Person) feather() { // 사람의 feathers 메서드 정의
	fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니다.")
}

type Quacker interface { // quack, feathers 메서드를 가지는 Quacker 인터페이스 정의
	quack()
	feather()
}

func inTheForest(q Quacker) {
	q.quack()   // Quacker 인터페이스로 quack 메서드 호출
	q.feather() // Quacker 인터페이스로 feathers 메서드 호출
}

func main() {
	var donald Duck // 오리 인스턴스 생성
	var john Person // 사람 인스턴스 생성

	inTheForest(donald) // 인터페이스를 통하여 오리의 quack, feather 메서드 호출
	inTheForest(john)   // 인터페이스를 통하여 사람의 quack, feather 메서드 호출

	//오리든 사람이든 상관없이 꽥(quack)과 깃털(feathers) 메서드만 가졌다면 모두 같은 인터페이스로 처리할 수 있습니다.


	//타입이 특정 인터페이스를 구현하는지 검사하려면 다음과 같이 사용합니다.
	//'interface{}(인스턴스).(인터페이스)'
	var donald2 Duck

	if v, ok := interface{}(donald2).(Quacker); ok { //Duck 타입의 인스턴스 donald2를 빈 인터페이스에 넣은 뒤 Quacker 인터페이스와 같은지 확인합니다.
		fmt.Println(v, ok) // {} true: 검사한 인스턴스는 구조체이며, 해당 인터페이스를 구현하고 있다.
	}
	//첫 번째 리턴값은 검사했던 인스턴스이며, 두 번째 리턴값은 인스턴스가 해당 인터페이스를 구현하고 있는지 여부입니다.
	//구현하고 있다면 true 그렇지 않으면 false입니다.
}
*/

// 빈 인터페이스 사용하기

/*
//인터페이스에 아무 메서드도 정의되어 있지 않으면 모든 타입을 저장할 수 있습니다.
type Any interface{} // 인터페이스에 메서드를 지정하지 않음

//빈 인터페이스 타입은 함수의 매개변수, 리턴값, 구조체의 필드로 사용할 수 있습니다.
func f2(arg Any) {   // 모든 타입을 저장할 수 있음
}
*/

/*
//이제 모든 타입을 받아서 내용을 출력하는 함수를 만들어보겠습니다.
package main

import (
	"fmt"
	"strconv"
)

//                      ↓ 빈 인터페이스를 사용하여 모든 타입을 받음
func formatString(arg interface{}) string {
	//       ↓ 인터페이스에 저장된 타입에 따라 case 실행
	switch arg.(type) {
	case int: // arg가 int형이라면
		i := arg.(int)         // int형으로 값을 가져옴
		return strconv.Itoa(i) // strconv.Itoa 함수를 사용하여 i를 문자열로 변환
	case float32: // arg가 float32형이라면
		f := arg.(float32)                                  // float32형으로 값을 가져옴
		return strconv.FormatFloat(float64(f), 'f', -1, 32) // strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case float64: // arg가 float64형이라면
		f := arg.(float64)                         // float64형으로 값을 가져옴
		return strconv.FormatFloat(f, 'f', -1, 64) // strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case string: // arg가 string이라면
		return arg.(string) // string이므로 그대로 리턴
	default:
		return "error"
	}

}

func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello, world!"))
}
*/

//일반 자료형뿐만 아니라 구조체 인스턴스 및 포인터도 빈 인터페이스로 넘길 수 있습니다.
package main

import (
	"fmt"
	"strconv"
)

type Person struct { // Person 구조체 정의
	name string
	age  int
}

func formatString2(arg interface{}) string {
	switch arg.(type) {
	case Person: // arg의 타입이 Person이라면
		name := arg.(Person).name // Person 타입으로 값을 가져옴
		age := strconv.Itoa(arg.(Person).age)
		return name + " " + age // 각 필드를 합쳐서 리턴
	case *Person: // arg의 타입이 Person 포인터라면
		name := arg.(*Person).name // *Person 타입으로 값을 가져옴
		age := strconv.Itoa((arg.(*Person).age))
		return name + " " + age // 각 필드를 합쳐서 리턴
	default:
		return "error"
	}
}

func main() {
	fmt.Println(formatString2(Person{"Maria", 20}))
	fmt.Println(formatString2(&Person{"John", 25}))

	andrew := Person{"andrew", 29}
	fmt.Println(formatString2(andrew))

	//인터페이스에 저장된 타입이 특정 타입인지 검사하려면 다음과 같이 사용합니다.
	var t interface{}
	t = Person{"Cen", 12}

	if v, ok := t.(Person); ok {
		//인터페이스.(타입) 형식입니다.
		//첫 번째 리턴값은 해당 타입으로 된 값이며
		//두 번째 리턴값은 타입이 맞는지 여부입니다.
		//타입이 일치하면 true 그렇지 않으면 false입니다.
		fmt.Println(v, ok) // {Cen 12} true
	}

}
