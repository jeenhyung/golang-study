//스택 2개로 큐 만들기

package main

import "fmt"

type IntBox []int
type Index uint

type Stack struct {
	values IntBox
	index  Index
}

func (s *Stack) push(item int) {
	s.values[s.index] = item
	s.index++
}
func (s *Stack) pop() int {
	if s.index > 0 {
		s.index--
	}
	item := s.values[s.index]
	s.values[s.index] = 0

	return item
}

func insert(s *Stack, item int) {
	s.push(item)
	fmt.Println("Succesed insert Item : ", item)
}

func remove(s1 *Stack) int {

	var s2 Stack
	s2.values = make(IntBox, s1.index+1)

	max := s1.index - 1
	var i Index
	//	fmt.Println("s1.index:", s1.index)
	for i = 0; i <= max; i++ {
		s2.push(s1.pop())
	}

	item := s2.pop()

	//	fmt.Println("s2.index:", s2.index)
	if s2.index > 0 {
		max = s2.index - 1
		for i = 0; i <= max; i++ {
			s1.push(s2.pop())
		}
	}
	fmt.Println("Succesed remove Item : ", item)
	return item
}

func StackView(s *Stack) {
	fmt.Println("")
	fmt.Println("[Stack View]")
	for index, value := range s.values {
		fmt.Printf("index:%d, value:%d \n", index, value)
	}
	fmt.Println("")
}
func main() {
	var stack Stack
	stack.values = make(IntBox, 5, 10)

	StackView(&stack)

	fmt.Println("[Insert Item]")
	insert(&stack, 1)
	insert(&stack, 2)
	insert(&stack, 3)

	StackView(&stack)
	_ = remove(&stack)
	StackView(&stack)
	_ = remove(&stack)
	StackView(&stack)
	_ = remove(&stack)
	StackView(&stack)
}
