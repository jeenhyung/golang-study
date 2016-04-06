// 문자열 입출력 함수 사용하기
//표준 출력(stdout), 표준 입력(stdin)뿐만 아니라 변수를 문자열로 만들거나 문자열에서 변수로 값을 가져올 수 있습니다.

/*
다음은 fmt 패키지에서 제공하는 문자열 입출력 함수입니다.

 - func Sprint(a …interface{}) string: 값을 그대로 문자열로 만듦
 - func Sprintln(a …interface{}) string: 값을 그대로 문자열로 만든 뒤 문자열 끝에 개행 문자(\n)를 붙임
 - func Sprintf(format string, a …interface{}) string: 형식을 지정하여 문자열을 만듦
 - func Sscan(str string, a …interface{}) (n int, err error): 공백, 개행 문자로 구분된 문자열에서 입력을 받음
 - func Sscanln(str string, a …interface{}) (n int, err error): 공백으로 구분된 문자열에서 입력을 받음
 - func Sscanf(str string, format string, a …interface{}) (n int, err error): 문자열에서 형식을 지정하여 입력을 받음
*/
//////////////////////////////////////////////////////////////////////////
/*
//먼저 변수 또는 값을 콘솔(터미널)로 출력하지 않고 문자열을 만들어보겠습니다.
package main

import "fmt"

func main() {
	var s1 string
	s1 = fmt.Sprint(1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만듦
	fmt.Print(s1)

	var s2 string
	s2 = fmt.Sprintln(2, 2.2, "Hello, world!2") // 값을 그대로 문자열로 만든 뒤 문자열 끝에 개행 문자를 붙임
	fmt.Print(s2)

	var s3 string
	s3 = fmt.Sprintf("%d %f %s\n", 3, 3.3, "Hello, world!3") // 형식을 지정하여 문자열로 만듦
	fmt.Print(s3)
}

*/
//////////////////////////////////////////////////////////////////////////

// 반대로 문자열에서 입력을 받을 수도 있습니다.
package main

import "fmt"

func main() {
	var num1 int
	var num2 float32
	var s string

	input1 := "1\n1.1\nHello"
	n, _ := fmt.Sscan(input1, &num1, &num2, &s) // 공백, 개행 문자로 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n)                    // 입력 개수: 3
	fmt.Println(num1, num2, s)                  // 1 1.1 Hello

	input2 := "2 2.2 Hello2"
	n, _ = fmt.Sscanln(input2, &num1, &num2, &s) // 공백으로 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n)                     //입력 개수: 3
	fmt.Println(num1, num2, s)                   // 2 2.2 Hello2

	input3 := "3,3.3,Hello3"
	n, _ = fmt.Sscanf(input3, "%d,%f,%s", &num1, &num2, &s) // 문자열에서 형식을 지정하여 입력을 받음
	fmt.Println("입력 개수:", n)                                //입력 개수: 3
	fmt.Println(num1, num2, s)                              // 3 3.3 Hello3
}
