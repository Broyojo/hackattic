package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func compress(str string) string {
	if len(str) == 0 {
		return str
	}

	curr := rune(str[0])
	count := 1
	var compressed string

	add := func() {
		if count > 2 {
			compressed += strconv.Itoa(count) + string(curr)
		} else {
			for i := 0; i < count; i++ {
				compressed += string(curr)
			}
		}
	}

	for i := 1; i < len(str); i++ {
		c := rune(str[i])
		if c != curr {
			add()
			curr = c
			count = 1
		} else {
			count++
		}
	}
	add()

	return compressed
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(compress(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
