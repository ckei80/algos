package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter word to reverse: ")
	text, _ := reader.ReadString('\n')
	reverseWords([]byte(text))
}

func reverseWords(s []byte) {
	var buffer bytes.Buffer

	maxLenght := len(s) - 1
	for _, value := range s {
		fmt.Println("moving: ", value)
		buffer.WriteByte(s[maxLenght])
		maxLenght--
	}
	fmt.Println(buffer.String())
}
