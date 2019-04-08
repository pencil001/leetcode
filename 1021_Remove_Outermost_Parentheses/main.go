package main

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))
}

func removeOuterParentheses(S string) string {
	posList := []Pos{}
	stack := &Stack{stack: []rune{}}
	curPos := Pos{b: -1, e: -1}
	for i, c := range S {
		if stack.Len() == 0 {
			stack.Push(c)
			curPos.b = i
		} else {
			p := stack.Peek()
			if p == '(' && c == ')' {
				stack.Pop()
			} else {
				stack.Push(c)
			}
			if stack.Len() == 0 {
				curPos.e = i
				posList = append(posList, curPos)
				curPos = Pos{b: -1, e: -1}
			}
		}
	}

	result := ""
	for _, p := range posList {
		result += string(S[p.b+1 : p.e])
	}
	return result
}

type Pos struct {
	b int
	e int
}

type Stack struct {
	stack []rune
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) Push(c rune) {
	s.stack = append(s.stack, c)
}

func (s *Stack) Pop() rune {
	if len(s.stack) == 0 {
		return 0
	}
	c := s.Peek()
	s.stack = s.stack[:len(s.stack)-1]
	return c
}

func (s *Stack) Peek() rune {
	if len(s.stack) == 0 {
		return 0
	}
	c := s.stack[len(s.stack)-1]
	return c
}
