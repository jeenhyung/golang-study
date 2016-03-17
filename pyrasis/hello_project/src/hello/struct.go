//구조체

package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

type Rectangle2 struct {
	width, height int // 자료형이 같은 필드는 한 줄로 나열
}

// 기본 사용법
/*
func main() {

	//일반 변수
	var rect Rectangle	// 구조체의 각 필드는 기본 값으로 초기화 됨
	var rect2 Rectangle = Rectangle{10, 20} 	// 구조체 인스턴스를 생성하면서 초기화
	rect3 := Rectangle{10, 20}	// var 키워드 생략. 구조체 인스턴스를 생성하면서 초기화
	rect4 := Rectangle{width: 10, height: 20} 	// 구조체 필드를 지정하여 값 초기화

	//포인터 변수
	var pRect *Rectangle	// 구조체 포인터 선언
	pRect = new(Rectangle)	// 구조체 포인터에 메모리 할당
	pRect2 := new(Rectangle)	// 구조체 포인터 선언과 동시에 메모리 할당


	//구조체 인스턴스 필드에 접근

	var rect4 Rectangle	// 구조체 인스턴스 생성
	var pRect3 *Rectangle	// 구조체 포인터 선언 후 메모리 할당

	rect4.width = 20	// 구조체 인스턴스의 필드에 접근할 때 .을 사용
	pRect3.width = 30	// 구조체 포인터의 필드에 접근할 때도 .을 사용

	fmt.Println(rect4)	// {20, 0}: 구조체 인스턴스 값
	fmt.Println(pRect3)	// &{30, 0}: 구조체 포인터므로 앞에 &가 붙음 (주소값이기 떄문)


}
*/

// 응용 사용법
/*
//구조체 생성자 패턴 활용하기

// new 함수로 구조체의 메모리를 할당하는 동시에 값을 초기화하는 방법은 없습니다.
// 하지만 다음과 같은 패턴을 사용하여 다른 언어의 생성자(Constructor)를 흉내낼 수 있습니다.
// Go 언어에서 지역 변수를 계속 참조하고 있다면, 스코프({ }로 묶인 영역)를 벗어나더라도 변수가 해제되지 않습니다.

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}	// 구조체 인스턴스를 생성한 뒤. 포인터를 리턴
}

func main() {
	//case1
	rect := NewRectangle(10, 20)
	//case2
	rect2 := &Rectangle{10, 20}

	fmt.Println(rect)	// &{10, 20}
	fmt.Println(rect2)	// &{10, 20}
}
*/

//함수에서 구조체 매개변수 사용하기
/*
//사각형 넓이 구하는 함수
func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}

func rectangleScaleA(rect *Rectangle, factor int) { // 매개변수로 구조체 포인터를 받음
	rect.width = rect.width * factor   // 포인터이므로 원래의 값이 변경됨
	rect.height = rect.height * factor // 포인터이므로 원래의 값이 변경됨
}

func rectangleScaleB(rect Rectangle, factor int) { // 매개변수로 구조체 인스턴스를 받음
	rect.width = rect.width * factor   // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
	rect.height = rect.height * factor // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
}

func main() {

	//사각형 넓이 구하기
	rect := Rectangle{30, 30}
	area := rectangleArea(&rect)
	fmt.Println(area)	// 900

	//call by reference
	rect1 := Rectangle{30, 30}
	rectangleScaleA(&rect1, 10) // 구조체의 포인터를 넘김
	fmt.Println(rect1)          // {300 300}: rect1에 바뀐 값이 들어감

	//call by value
	rect2 := Rectangle{30, 30}
	rectangleScaleB(rect2, 10) // 구조체 인스턴스를 그대로 넘김
	fmt.Println(rect2)         // {30 30}: rect2는 값이 바뀌지 않음
}
*/

//구조체에 메서드 연결하기
/*
//Go 언어에는 클래스가 없는 대신 구조체에 메서드를 연결할 수 있습니다.
// 'func (리시버명 *구조체_타입) 함수명() 리턴값_자료형 { }'
//리시버명은 호출 하는 쪽의 변수명과 달라도 된다.
//구조체 타입에 의해 구조체-함수가 연결된다.

//             ↓ 리시버 변수 정의(연결할 구조체 정의)
func (rect *Rectangle) area() int {
	return rect.width * rect.height // 리시버 변수를 사용하여 현재 인스턴스에 접근 할 수 있음
}

//           ↓ 포인터 방식
func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor   // 포인터이므로 원래의 값이 변경됨
	rect.height = rect.height * factor // 포인터이므로 원래의 값이 변경됨
}

//          ↓ 일반 구조체 방식
func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor   // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
	rect.height = rect.height * factor // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
}

func (_ Rectangle) dummy() { // _로 리시버 변수 생략
	fmt.Println("dummy")
}

func main() {
	//기본 예
	rect := Rectangle{10, 20}
	fmt.Println(rect.area())

	//함수에 구조체 형태의 매개변수를 넘기는 방식과 마찬가지로
	//리시버 변수를 받는 방법도 '포인터'와 '일반 구조체' 방식이 있습니다.

	//포인터 방식 호출 예
	rect1 := Rectangle{30, 30}
	rect1.scaleA(10)
	fmt.Println(rect1) // {300 300}: rect1에 바뀐 값이 들어감

	// 일반 구조체 방식 호출 예
	rect2 := Rectangle{30, 30}
	rect2.scaleB(10)
	fmt.Println(rect2) // {30 30}: rect2는 값이 바뀌지 않음

	//리시버 변수를 사용하지 않는다면 _ (밑줄 문자)로 변수를 생략할 수 있습니다.
	var r Rectangle
	r.dummy() // dummy

}
*/

// 구조체 임베딩 사용하기

//Go 언어는 클래스를 제공하지 않으므로 상속 또한 제공하지 않습니다.
//하지만 구조체에서 임베딩(Embedding)을 사용하면 상속과 같은 효과를 낼 수 있습니다.

type Person struct { // 사람 구조체 정의
	name string
	age  int
}

func (p *Person) greeting() { // 인사(greeting) 함수 작성
	fmt.Println("Hello~")
}

//case1
type Student struct {
	person Person // 학생 구조체 안에 사람 구조체를 필드로 가지고 있음. Has-a 관계 => "학생은 사람을 가진다."
	school string
	grade  int
}

//case2
type Student2 struct {
	Person // 익명 필드. 필드명 없이 타입만 선언하면 포함(Is-a) 관계가 됨 => “학생은 사람이다.”
	school string
	grade  int
}

//오버라이딩
func (p *Student2) greeting() { // 이름이 같은 메서드를 정의하면 오버라이딩됨
	fmt.Println("Hello Students~")
}

func main() {

	//Has-a 관계 예
	var student Student
	student.person.greeting() // Hello~

	//이번에는 Student 구조체에 Person 구조체를 임베딩합니다.
	//Is-a 관계 예
	var student2 Student2
	student2.Person.greeting() // Hello~
	student2.greeting()        // Hello Students~

}

//구조체 임베딩을 사용하면 다른 언어의 '상속'과 동일한 형태가 됩니다.
//물론 구조체 안에 여러 개의 구조체를 임베딩하면 다중 상속도 구현할 수 있습니다.
//하지만 Go 언어에서는 복잡한 다중 상속 대신 인터페이스를 주로 활용합니다.
//인터페이스에 대해서는 뒤에서 자세히 설명하겠습니다.
