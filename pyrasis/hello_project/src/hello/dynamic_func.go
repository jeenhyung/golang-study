// 동적으로 함수 생성하기

//이번에는 리플렉션을 사용하여 동적으로 함수를 만들어내는 방법을 알아보겠습니다.

/*
//먼저 reflect.MakeFunc 함수를 사용하는 방법입니다.
package main

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value { // 매개변수와 리턴값은 반드시 []reflect.Value를 사용
	fmt.Println("Hello, world!")
	return nil
}

func main() {
	var hello func() // 함수를 담을 변수 선언

	fn := reflect.ValueOf(&hello).Elem() // hello의 주소를 넘긴 뒤 Elem으로 값 정보를 가져옴

	v := reflect.MakeFunc(fn.Type(), h) // h의 함수 정보를 생성

	fn.Set(v) // hello의 값 정보인 fn에 h의 함수 정보 v를 설정하여 함수를 연결

	hello()
	//hello 함수는 h 함수를 이용하여 동적으로 생성된 함수입니다. 하지만 Hello, world!를 출력하기에는 복잡하기만 합니다.
}
*/

//동적 함수 생성을 좀더 제대로 활용하기 위해 함수 하나로 정수, 실수, 문자열 더하기를 모두 처리하는 함수를 생성해보겠습니다.

package main

import (
	"fmt"
	"reflect"
)

//두 값을 더하는 sum 함수를 구현합니다.
func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("타입이 다릅니다.")
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

//이렇게 하면 매개변수, 리턴값 자료형의 형태가 다양한 함수를 동적으로 연결시킬 수 있습니다.
func makeSum(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), sum)

	fn.Set(v)
}

func main() {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	makeSum(&intSum)
	makeSum(&floatSum)
	makeSum(&stringSum)

	fmt.Println(intSum(1, 2))                   // 3
	fmt.Println(floatSum(2.1, 3.5))             // 5.599999904632568
	fmt.Println(stringSum("Hello, ", "world!")) // Hello, world!
}
