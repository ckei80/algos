package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []int

func main() {
	tokens := []string{"1", "2", "+", "3", "*"}
	operators := "+-/*"
}

func EvalRPN() {

}

func (s stack) Push(v int) {
	return append(s, v)
}

func (stack s) Pop(stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}
