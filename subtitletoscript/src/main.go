package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("gameofthrones-1-1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data[0])
	fmt.Println(data[1])

	data_s := strings.Replace(string(data), "<br>", " ", -1)   // 줄바꿈 => 뛰어쓰기로 변경
	data_s = strings.Replace(string(data_s), "&nbsp;", "", -1) // &nbsp; => 빈값으로 변경
	data_s = strings.Replace(string(data_s), "<P Class=KRCC>", "", -1)
	data_s = strings.Replace(string(data_s), "<P Class=ENCC>", "", -1)

	data = []byte(data_s)

	b_del := false
	for i := 0; i < len(data); i++ {
		if data[i] == 60 && b_del == false {
			b_del = true
			data[i] = 0
		} else if data[i] == 62 && b_del == true {
			data[i] = 0
			b_del = false
		}
		if b_del == true {
			data[i] = 0
		}
	}
	err = ioutil.WriteFile("gameofthrones-1-1copy.txt", []byte(string(data)), os.FileMode(644))
	if err != nil {
		fmt.Println(err)
		return
	}
}
