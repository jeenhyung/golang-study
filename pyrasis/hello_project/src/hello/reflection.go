///////////////////////////////////////////////////////
// 리플렉션 사용하기
///////////////////////////////////////////////////////

//리플렉션은 실행 시점(Runtime, 런타임)에 인터페이스나 구조체 등의 타입 정보를 얻어내거나 결정하는 기능입니다.

/*
//간단하게 변수와 구조체의 타입을 표시해보겠습니다.
package main

import (
	"fmt"
	"reflect"
)

type Data struct { // 구조체 정의
	a, b int
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num)) // int : reflect.TypeOf로 자료형 이름 출력

	var s string = "Hello, world!"
	fmt.Println(reflect.TypeOf(s)) // string : reflect.TypeOf로 자로형 이름 출력

	var f float32 = 1.3
	fmt.Println(reflect.TypeOf(f)) // float32: reflect.TypeOf로 자료형 이름 출력

	var data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data)) // main.Data: reflect.TypeOf로 구조체 이름 출력

	//int, string, float32 형 변수의 자료형이 출력됩니다.
	//마찬가지로 구조체도 타입을 알아낼 수 있는데 Data 구조체는 main 패키지 안에 속해있기 때문에 main.Data로 나옵니다.
}
*/
/*
//리플렉션으로 변수의 타입뿐만 아니라 값에 대한 상세한 정보도 얻어올 수 있습니다.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 1.3
	t := reflect.TypeOf(f)  // f의 타입 정보를 t에 저장
	v := reflect.ValueOf(f) // f의 값 정보를 v에 저장

	fmt.Println(t.Name())                    // float64: 자로형 이름 출력
	fmt.Println(t.Size())                    // 8: 자료형 크기 출력
	fmt.Println(t.Kind() == reflect.Float64) // true: 자로형 종류를 알아내어 reflect.Float64와 비교
	fmt.Println(t.Kind() == reflect.Int64)   // false: 자료형 종류를 알아내어 reflect.Int64와 비교

	//reflect.ValueOf 함수로 float64 변수의 값 정보 reflect.Value를 얻어오면 타입 정보, 타입 종류, 변수에 저장된 값을 알 수 있습니다.
	fmt.Println(v.Type())                    // float64: 값이 담긴 변수의 자료형 이름 출력
	fmt.Println(v.Kind() == reflect.Float64) // true: 값이 담긴 변수의 자료형 종류를 알아내어 reflect.Float64와 비교
	fmt.Println(v.Kind() == reflect.Int64)   // false: 값이 담긴 변수의 자료형 종류를 알아내어 reflect.Int64와 비교

	//변수가 float64라면 v.Float(), int라면 v.Int(), string이라면 v.String()과 같이 각 타입에 맞는 함수를 사용하면 변수에 저장된 값을 가져올 수 있습니다.
	fmt.Println(v.Float()) // 1.3: 값을 실수형으로 출력
}
*/

///////////////////////////////////////////////////////////////////////
// 구조체 태그 가져오기
/*
//다음과 같이 리플렉션으로 구조체의 태그도 가져올 수 있습니다.
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	//구조체 필드의 태그는 `태그명:"내용"` 형식으로 지정합니다.
	//태그는 문자열 형태이며 문자열 안에 " " (따옴표)가 포함되므로 ` ` (백쿼트)로 감싸줍니다.
	//여러 개를 지정할 때는 공백으로 구분해줍니다.
	name string 'tag1:"이름" tag2:"Name"'		// 구조체에 태그 설정
	age  int    'tag1:"나이" tag2:"Age"'		// 구조체에 태그 설정
}

func main() {
	p := Person()

	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))	// true 이름 Name

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))	// true 나이 Age

}
*/

////////////////////////////////////////////////////////////////////

// 포인터와 인터페이스의 값 가져오기

//다음은 일반 포인터와 인터페이스의 값을 가져오는 방법입니다.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a *int = new(int)
	*a = 1

	fmt.Println(reflect.TypeOf(a))  // *int
	fmt.Println(reflect.ValueOf(a)) // <*int Value>
	//int 포인터 a의 값 정보에서 바로 Int 함수로 값을 가져오려면 런타임 에러가 발생합니다.
	//따라서 Elem 함수로 포인터의 메모리에 저장된 실제 값 정보를 가져와야 합니다.
	//fmt.Println(reflect.ValueOf(a).Int())        // 런타임 에러
	fmt.Println(reflect.ValueOf(a).Elem())       // 1: <int Value>
	fmt.Println(reflect.ValueOf(a).Elem().Int()) // 1

	fmt.Println()

	var b interface{}
	b = 1

	//빈 인터페이스 b에 1을 대입하면 타입 정보는 int이고 값 정보는 int입니다.
	fmt.Println(reflect.TypeOf(b))        // int
	fmt.Println(reflect.ValueOf(b))       // 1: <int Value>
	fmt.Println(reflect.ValueOf(b).Int()) // 1
	//fmt.Println(reflect.ValueOf(b).Elem()) // 런타임 에러

}
