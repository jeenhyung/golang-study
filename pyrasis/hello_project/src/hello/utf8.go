// UTF-8 함수 사용하기
//UTF-8은 유니코드를 저장하거나 전송할 때 사용하는 인코딩 방식 중의 하나입니다(UTF-8 이외에도 UTF-7, UTF-16, UTF-32 등 다양한 방식이 있습니다). Go 언어에서는 UTF-8을 주로 사용하므로 UTF-8에 대해 자세히 알아보겠습니다.
/*
다음은 unicode/utf8 패키지에서 제공하는 함수입니다.
 - func RuneLen(r rune) int: 문자의 바이트 수를 구함
 - func RuneCountInString(s string) (n int): 문자열의 실제 길이를 구함
 - func DecodeRune(p []byte) (r rune, size int): byte 슬라이스에서 첫 글자를 디코딩함
 - func DecodeLastRune(p []byte) (r rune, size int): byte 슬라이스에서 마지막 글자를 디코딩함
 - func DecodeRuneInString(s string) (r rune, size int): 문자열에서 첫 글자를 디코딩함
 - func DecodeLastRuneInString(s string) (r rune, size int): 문자열에서 마지막 글자를 디코딩함
 - func Valid(p []byte) bool: byte 슬라이스가 UTF-8이 맞는지 확인
 - func ValidRune(r rune) bool: rune 변수에 저장된 값이 UTF-8이 맞는지 확인
 - func ValidString(s string) bool: 문자열이 UTF-8이 맞는지 확인
*/
//UTF-8은 가변 길이 문자 인코딩 방식이라 문자를 저장할 때 1바이트에서 4바이트까지 사용하며 한글은 3바이트로 저장합니다.
//
////////////////////////////////////////////////////////////////////////////
/*
//먼저 한글 글자 하나의 길이(바이트 수)를 구해보겠습니다.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string = "한"
	fmt.Println(len(s)) // 3: 한글은 3바이트로 저장하므로 3

	var r rune = '한'
	fmt.Println(utf8.RuneLen(r)) // 3: 한글은 3바이트로 저장하므로 3

	//한글 문자열에서 바이트 수가 아닌 실제 길이(글자 개수)를 구해보겠습니다.
	var s1 string = "안녕하세요"
	fmt.Println(utf8.RuneCountInString(s1))	// 5: "안녕하세요"의 실제 길이는 5
}
*/
//////////////////////////////////////////////////////////////////////////
/*
// 이번에는 한글 문자열에서 글자를 디코딩해보겠습니다.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("안녕하세요")

	r, size := utf8.DecodeRune(b)
	fmt.Printf("%c %d\n", r, size) // 안 3: "안녕하세요"의 첫 글자를 디코딩하여 '안', 바이트 수 3

	r, size = utf8.DecodeRune(b[3:]) // '안'의 길이가 3이므로 인덱스 3부터 부분 슬라이스를 만들면 "녕하세요"가 됨
	fmt.Printf("%c %d\n", r, size)   // 녕 3: "녕하세요"를 첫 글자를 디코딩하여 '녕', 바이트 수 3

	r, size = utf8.DecodeLastRune(b)
	fmt.Printf("%c %d\n", r, size) // 요 3: "안녕하세요"의 마지막 글자를 디코딩하여 '요', 바이트 수 3

	// '요'의 길이가 3이므로 문자열 길이-3을 하여 부분 슬라이스를 만들면 "안녕하세"가 됨
	r, size = utf8.DecodeLastRune(b[:len(b)-3])
	fmt.Printf("%c %d\n", r, size) // 세 3: "안녕하세"의 마지막 글자를 디코딩하여 '세', 바이트 수 3

	r, size = utf8.DecodeRune(b[:len(b)-3])
	fmt.Printf("%c %d\n", r, size) // 안 3: "안녕하세"의 첫 글자를 디코딩하여 '안', 바이트 수 3
}
*/
///////////////////////////////////////////////////////////////////////////

// 그러면 문자열의 첫 글자와 마지막 글자를 구하려면 어떻게 해야 할까요?
// 다음은 영문 문자열의 첫 글자와 마지막 글자를 구합니다
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, world!"

	fmt.Printf("%c\n", s[0])        // H: 인덱스 0이 첫 번째 글자
	fmt.Printf("%c\n", s[len(s)-1]) // !: 문자열 길이에서 1을 뺀 인덱스가 마지막 글자

	//한글 문자열은 UTF-8에서 3바이트로 저장되므로 인덱스로 접근하면 한글이 정상적으로 출력되지 않습니다.
	//따라서 한글 문자열의 첫 글자와 마지막 글자를 구하려면 다음과 같이 utf8.DecodeRuneInString, utf8.DecodeLastRuneInString 함수를 사용하면 됩니다.
	s = "안녕하세요"

	r, _ := utf8.DecodeRuneInString(s) // UTF-8 문자열의 첫 글자와 바이트 수를 리턴
	fmt.Printf("%c\n", r)              // 안: 문자열의 첫 글자

	r, _ = utf8.DecodeLastRuneInString(s) // UTF-8 문자열의 마지막 글자와 바이트 수를 리턴
	fmt.Printf("%c\n", r)                 // 요: 문자열의 마지막 글자

	//값이나 문자열이 UTF-8이 맞는지 확인하는 방법은 다음과 같습니다.
	var b1 []byte = []byte("안녕하세요")
	fmt.Println(utf8.Valid(b1)) // true: "안녕하세요"는 UTF-8이 맞으므로 true
	var b2 []byte = []byte{0xff, 0xf1, 0xc2}
	fmt.Println(utf8.Valid(b2)) // false: 0xff 0xf1 0xc1은 UTF-8이 아니므로 false
	var b3 []byte = []byte("hello")
	fmt.Println(utf8.Valid(b3)) // true: "hello"는 UTF-8이 맞으므로 true

	var r1 rune = '한'
	fmt.Println(utf8.ValidRune(r1)) // true: '한'은 UTF-8이 맞으므로 true
	var r2 rune = 0x11111111
	fmt.Println(utf8.ValidRune(r2)) // false: 0x11111111은 UTF-8이 아니므로 false
	var r3 rune = 'a'
	fmt.Println(utf8.ValidRune(r3)) // true: 'a'는 UTF-8이 맞으므로 true

	var s1 string = "한글"
	fmt.Println(utf8.ValidString(s1)) // true: "한글"은 UTF-8이 맞으므로 true
	var s2 string = string([]byte{0xff, 0xf1, 0xc1})
	fmt.Println(utf8.ValidString(s2)) // false: 0xff, 0xf1, 0xc1 UTF-8이 아니므로 false
}
