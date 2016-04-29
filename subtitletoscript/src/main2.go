package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
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
	// < > 속 메타 설정 문자 삭제
	for i := 0; i < len(res); i++ {
		if res[i] == 60 && b_del == false {
			b_del = true
			res[i] = 0
		} else if res[i] == 62 && b_del == true {
			res[i] = 0
			b_del = false
		}
		if b_del == true {
			res[i] = 0
		}
	}
	err = ioutil.WriteFile("gameofthrones-1-1-bymain2.txt", []byte(string(res)), os.FileMode(644))
	if err != nil {
		fmt.Println(err)
		return
	}
}
