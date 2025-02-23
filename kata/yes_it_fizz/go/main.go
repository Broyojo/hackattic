package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	parts := strings.Split(strings.Trim(text, "\n"), " ")
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	for i := n; i <= m; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
