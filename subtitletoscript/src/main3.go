package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	data, err := ioutil.ReadFile("gameofthrones-1-1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("< byte code: ", data[0])
	fmt.Println("> byte code: ", data[1])

	data_s := strings.Replace(string(data), "<br>", " ", -1)   // 줄바꿈 => 뛰어쓰기로 변경
	data_s = strings.Replace(string(data_s), "&nbsp;", "", -1) // &nbsp; => 빈값으로 변경
	data_s = strings.Replace(string(data_s), "<P Class=KRCC>", "", -1)
	data_s = strings.Replace(string(data_s), "<P Class=ENCC>", "", -1)

	// data = []byte(data_s)

	re := regexp.MustCompile("(?m)[\r\n]+^.*SYNC.*$")
	data_s = re.ReplaceAllString(data_s, "") // "SYNC"를 포함한 행 삭제
	re = regexp.MustCompile("(?m)[\r\n]+^.*TITLE.*$")
	res := []byte(re.ReplaceAllString(data_s, "")) // "TITLE"를 포함한 행 삭제

	b_del := false
	for i := 0; i < len(res); i++ {
		if res[i] == 60 && b_del == false { // '<'부터 삭제
			b_del = true
			res[i] = 0
		} else if res[i] == 62 && b_del == true { // '>'까지 삭제
			res[i] = 0
			b_del = false
		}
		if b_del == true {
			res[i] = 0
		}
	}

	f1 := func(r rune) bool {
		return !unicode.Is(unicode.Hangul, r) // r이 한글 유니코드이면 false를 리턴
	}
	f2 := func(r rune) bool {
		return !unicode.Is(unicode.Latin, r) // r이 한글 유니코드이면 true를 리턴
	}

	//한글자막 영어자막 분리
	script_kr := strings.TrimFunc(string(res), f1)
	script_eng := strings.TrimFunc(string(res), f2)
	// fmt.Println(script_kr)
	fmt.Println(script_eng)

	// script_kr := strings.FieldsFunc(string(res), f)
	// stringByte := "\x00" + strings.Join(script_kr, "\x20\x00") // []string 을 []byte로 전환하기전 설정

	//한글자막 영어자막 한줄씩 썪기

	//파일 생성
	err = ioutil.WriteFile("gameofthrones-1-1-bymain3-kr.txt", []byte(script_kr), os.FileMode(644))
	if err != nil {
		fmt.Println(err)
		return
	}
	//파일 생성
	err = ioutil.WriteFile("gameofthrones-1-1-bymain3-eng.txt", []byte(script_eng), os.FileMode(644))
	if err != nil {
		fmt.Println(err)
		return
	}
}
