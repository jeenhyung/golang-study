// 문자열 처리하기
// Go 언어는 기본 라이브러리에서 다양한 문자열 처리 함수를 제공하므로 이 함수들을 조합하면 코드를 간단하게 유지할 수 있고, 관련 기능을 다시 구현하지 않아도 됩니다.

// 문자열 검색하기
//문자열을 처리할 때 주로 사용하는 기능은 문자열 검색입니다.
/*
다음은 strings 패키지에서 제공하는 문자열 검색 함수입니다.

 - func Contains(s, substr string) bool: 문자열이 포함되어 있는지 검색
 - func ContainsAny(s, chars string) bool: 특정 문자가 하나라도 포함되어 있는지 검색
 - func ContainsRune(s string, r rune) bool: rune 자료형으로 검색
 - func Count(s, sep string) int: 문자열이 몇 번 나오는지 구함
 - func HasPrefix(s, prefix string) bool: 문자열이 접두사인지 판단
 - func HasSuffix(s, suffix string) bool: 문자열이 접미사인지 판단
 - func Index(s, sep string) int: 특정 문자열의 위치를 구함
 - func IndexAny(s, chars string) int: 가장 먼저 나오는 문자의 위치를 구함
 - func IndexByte(s string, c byte) int: byte 자료형으로 위치를 구함
 - func IndexRune(s string, r rune) int: rune 자료형으로 위치를 구함
 - func IndexFunc(s string, f func(rune) bool) int: 검색 함수를 정의하여 위치를 구함
 - func LastIndex(s, sep string) int: 가장 마지막에 나오는 특정 문자열의 위치를 구함
 - func LastIndexAny(s, chars string) int: 가장 마지막에 나오는 문자의 위치를 구함
 - func LastIndexFunc(s string, f func(rune) bool) int: 검색 함수를 정의하여 위치를 구함
*/

// 먼저 Hello, world! 문자열에서 특정 문자열 및 문자를 찾아보겠습니다.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, world!", "wo"))  // true
	fmt.Println(strings.Contains("Hello, world!", "w o")) // false
	fmt.Println(strings.Contains("Hello, world!", "ow"))  // false

	fmt.Println(strings.ContainsAny("Hello, world!", "wo"))  // true
	fmt.Println(strings.ContainsAny("Hello, world!", "w o")) // true
	fmt.Println(strings.ContainsAny("Hello, world!", "ow"))  // true

	fmt.Println(strings.Count("Hello Helium", "He")) // 2

	var r rune
	r = '하'
	fmt.Println(strings.ContainsRune("안녕하세요", r)) // true

	fmt.Println(strings.HasPrefix("Hello, world!", "He"))   // true
	fmt.Println(strings.HasSuffix("Hello, world!", "rld!")) // true
}
