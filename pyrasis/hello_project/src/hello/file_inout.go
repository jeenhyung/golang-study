// 파일 입출력 함수 사용하기
//값을 파일로 저장하거나 파일에서 변수로 값을 가져올 수 있습니다.
/*
다음은 os 패키지에서 제공하는 파일 처리 함수입니다.

 - func Create(name string) (file *File, err error): 기존 파일을 열거나 새 파일을 생성
 - func Open(name string) (file *File, err error): 기존 파일을 열기
 - func (f *File) Close() error: 열린 파일을 닫음
*/

/*
다음은 fmt 패키지에서 제공하는 입출력 함수이며 함수를 사용하기 전에 os 패키지의 파일 처리 함수로 파일을 생성하거나 열어야 합니다.

 - func Fprint(w io.Writer, a …interface{}) (n int, err error): 값을 그대로 문자열로 만든 뒤 파일에 저장
 - func Fprintln(w io.Writer, a …interface{}) (n int, err error): 값을 그대로 문자열로 만든 뒤 문자열 끝에 개행 문자(\n)를 붙이고 파일에 저장
 - func Fprintf(w io.Writer, format string, a …interface{}) (n int, err error): 형식을 지정하여 파일에 저장
 - func Fscan(r io.Reader, a …interface{}) (n int, err error): 파일을 읽은 뒤 공백, 개행 문자로 구분된 문자열에서 입력을 받음
 - func Fscanln(r io.Reader, a …interface{}) (n int, err error): 파일을 읽은 뒤 공백으로 구분된 문자열에서 입력을 받음
 - func Fscanf(r io.Reader, format string, a …interface{}) (n int, err error): 파일을 읽은 뒤 문자열에서 형식을 지정하여 입력을 받음
*/
////////////////////////////////////////////////////////////////////
/*
//먼저 값을 파일로 저장해보겠습니다
package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Create("hello1.txt")        //hello1.txt 파일 생성
	defer file1.Close()                        // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprint(file1, 1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만든 뒤 파일에 저장

	file2, _ := os.Create("hello2.txt")
	defer file2.Close()
	fmt.Fprintln(file2, 2, 2.2, "Hello, world!2") // 값을 그대로 문자열로 만든 뒤 문자열 끝에 개행 문자를 붙이고 파일에 저장

	file3, _ := os.Create("hello3.txt")
	defer file3.Close()
	fmt.Fprintf(file3, "%d,%f,%s", 3, 3.3, "Hello, world!3") // 형식을 지정하여 파일에 저장
}
*/
///////////////////////////////////////////////////////////////////////////

//반대로 파일에서 입력을 받아보겠습니다. 파일은 방금 만든 파일(hello1.txt, hello2.txt, hello3.txt)을 그대로 사용합니다.
package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 int
	var num2 float32
	var s string

	file1, _ := os.Open("hello1.txt")
	defer file1.Close()
	n, _ := fmt.Fscan(file1, &num1, &num2, &s) // 파일을 읽은 뒤 공백, 개행 문자로 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n)                   // 입력 개수: 3
	fmt.Println(num1, num2, s)

	file2, _ := os.Open("hello2.txt")
	defer file2.Close()
	n, _ = fmt.Fscanln(file2, &num1, &num2, &s) // 파일을 읽은 뒤 공백으로 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n)
	fmt.Println(num1, num2, s)

	file3, _ := os.Open("hello3.txt")
	defer file3.Close()
	n, _ = fmt.Fscanf(file3, "%d,%f,%s", &num1, &num2, &s) // 파일을 읽은 뒤 문자열에서 형식을 지정하여 입력을 받음
	fmt.Println("입력 개수:", n)
	fmt.Println(num1, num2, s)

}
