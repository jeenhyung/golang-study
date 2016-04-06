//여기서 디렉터리의 이름은 패키지 이름과 동일하게 만듭니다.
//즉, calc라는 패키지가 있으면 디렉터리는 GOPATH/src/calc가 됩니다.
//.go 소스 파일의 이름은 패키지 이름과 같지 않아도 되며 각자 상황에 맞게 적절히 짓습니다.

//소스 파일의 첫 줄에서 package calc로 설정하여 현재 파일이 calc 패키지에 포함된다는 것을 알려줍니다.
package calc

//패키지 안에서 함수, 변수, 상수의 이름을 정하는 방법은 두 가지가 있습니다.
// - 첫 글자를 영문 소문자로 지정하면 패키지 안에서만 사용할 수 있습니다. 즉 외부에서 사용할 수 없습니다. 예) sum, max, hello
// - 첫 글자를 영문 대문자로 지정하면 외부에서 사용할 수 있습니다. 예) Sum, Max, Hello

//두 정수를 더함
func Sum(a, b int) int { // 외부에서 사용할 수 있도록 함수의 첫 글자는 영문 대문자로
	return a + b
}

// GOPATH/src/hello/hello.go 에서 사용