// 유니코드 함수 사용하기
/*
다음은 Unicode 패키지에서 제공하는 함수입니다.

 - func Is(rangeTab *RangeTable, r rune) bool: 문자가 지정한 범위 테이블에 포함되는지 확인
 - func In(r rune, ranges …*RangeTable) bool: 문자가 여러 범위 테이블 중에 포함되는지 확인
 - func IsGraphic(r rune) bool: 값이 화면에 표시될 수 있는지 확인
 - func IsLetter(r rune) bool: 값이 문자인지 확인
 - func IsDigit(r rune) bool: 값이 숫자인지 확인
 - func IsControl(r rune) bool: 값이 제어 문자인지 확인
 - func IsMark(r rune) bool: 값이 마크인지 확인
 - func IsPrint(r rune) bool: 값이 Go 언어에서 출력할 수 있는지 확인
 - func IsPunct(r rune) bool: 값이 문장 부호인지 확인
 - func IsSpace(r rune) bool: 값이 공백인지 확인
 - func IsSymbol(r rune) bool: 값이 심볼인지 확인
 - func IsUpper(r rune) bool: 값이 대문자인지 확인
 - func IsLower(r rune) bool: 값이 소문자인지 확인


다음은 대표적인 범위 테이블(unicode.RangeTable)입니다.

 - unicode.Latin: 라틴 문자, 로마자, 영문자
 - unicode.Hangul: 한글
 - unicode.Han: 한자
 - unicode.Hiragana, unicode.Katakana: 일본어 히라가나, 카타카나

*/

/*
//먼저 문자가 유니코드인지 확인하는 방법입니다.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	var r1 rune = '한'
	fmt.Println(unicode.Is(unicode.Hangul, r1)) // true: r1은 한글이므로 true
	fmt.Println(unicode.Is(unicode.Latin, r1))  // false: r1은 라틴 문자가 아니므로 false

	var r2 rune = '鎭'
	fmt.Println(unicode.Is(unicode.Han, r2))    // true: r2는 한자이므로 true
	fmt.Println(unicode.Is(unicode.Hangul, r2)) // false: r2는 한글이 아니므로 false

	var r3 rune = 'a'
	fmt.Println(unicode.Is(unicode.Latin, r3))  // true: r3은 라틴 문자이므로 true
	fmt.Println(unicode.Is(unicode.Hangul, r3)) // false: r3은 한글이 아니므로 false

	//unicode.In 함수는 문자가 여러 범위 테이블 중에 포함되는지 확인할 수 있습니다.
	fmt.Println(unicode.In(r1, unicode.Latin, unicode.Han, unicode.Hangul)) // true: r1은 한글이므로 true
}
*/

//이번에는 유니코드 문자의 특성을 확인해보겠습니다.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(unicode.IsGraphic('1'))  // true: 1은 화면에 표시되는 숫자이므로 true
	fmt.Println(unicode.IsGraphic('a'))  // true: a는 화면에 표시되는 문자이므로 true
	fmt.Println(unicode.IsGraphic('한'))  // true: '한'은 화면에 표시되는 문자이므로 true
	fmt.Println(unicode.IsGraphic('漢'))  // true: '漢'은 화면에 표시되는 문자이므로 true
	fmt.Println(unicode.IsGraphic('\n')) // false: \n 화면에 표시되는 문자가 아니므로 false

	fmt.Println(unicode.IsLetter('a')) // true: a는 문자이므로 true
	fmt.Println(unicode.IsLetter('1')) // false: 1은 문자가 아니므로 false

	fmt.Println(unicode.IsDigit('1'))     // true: 1은 숫자이므로 true
	fmt.Println(unicode.IsControl('\n'))  // true: \n은 제어 문자이므로 true
	fmt.Println(unicode.IsMark('\u17c9')) // true: \u17c9는 마크이므로 true

	fmt.Println(unicode.IsPrint('1')) // true: 1은 Go 언어에서 표시할 수 있으므로 true
	fmt.Println(unicode.IsPunct('.')) // true: .은 문장 부호이므로 true

	fmt.Println(unicode.IsSpace(' '))  // true: ' '는 공백이므로 true
	fmt.Println(unicode.IsSymbol('♥')) // true: ♥는 심볼이므로 true

	fmt.Println(unicode.IsUpper('A')) // true: A는 대문자이므로 true
	fmt.Println(unicode.IsLower('a')) // true: a는 소문자이므로 true
}
