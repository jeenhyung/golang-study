// 입력 함수 사용하기

/*
다음은 fmt 패키지에서 제공하는 표준 입력 함수입니다.
 - func Scan(a …interface{}) (n int, err error): 콘솔에서 공백, 새 줄로 구분하여 입력을 받음
 - func Scanln(a …interface{}) (n int, err error): 콘솔에서 공백으로 구분하여 입력을 받음
 - func Scanf(format string, a …interface{}) (n int, err error): 콘솔에서 형식을 지정하여 입력을 받음
*/

//이번에는 콘솔에서 입력을 받아보겠습니다.
package main

import "fmt"

func main() {
	var s1, s2 string
	n, _ := fmt.Scan(&s1, &s2) // fmt.Scan 함수의 두 번째 리턴값은 생략
	fmt.Println("입력 개수:", n)
	fmt.Println(s1, s2)

	var num1, num2 int

	n, _ = fmt.Scanf("%d,%d", &num1, &num2) // 정수형으로 형식을 지정하여 입력을 받음
	//포맷을 설정할 때 %d,%d로 했으므로 값을 공백으로 구분하지 않고 콤마로 구분하여 입력을 받게 됩니다.
	fmt.Println("입력 개수:", n)
	fmt.Println(num1, num2)

}
