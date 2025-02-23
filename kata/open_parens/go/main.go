package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(parens string) bool {
	var count int
	for _, c := range parens {
		switch c {
		case '(':
			count++
		case ')':
			if count <= 0 {
				return false
			}
			count--
		default:
			panic(fmt.Sprintf("unknown character: %v\n", c))
		}
	}
	return count == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if check(line) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
