package main

import (
	"bufio"
	"fmt"
	stack "github.com/sikender/go-stack"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	st := new(stack.Stack)

	for {
		var token string
		s, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		for _, char := range s {
			switch {
			case char >= '0' && char <= '9':
				token = token + string(char)
			case char == ' ':
				r, _ := strconv.Atoi(token)
				st.Push(r)
				token = ""
			case char == '+':
				fmt.Printf("%v\n", st.Pop()+st.Pop())
			case char == '-':
				a, b := st.Pop(), st.Pop()
				fmt.Printf("%v\n", b-a)
			case char == '*':
				fmt.Printf("%v\n", st.Pop()*st.Pop())
			case char == 'q':
				return
			default:
			}
		}
	}
}
