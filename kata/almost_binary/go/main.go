package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.ReplaceAll(text, "#", "1")
		text = strings.ReplaceAll(text, ".", "0")
		num, err := strconv.ParseUint(text, 2, 16)
		if err != nil {
			panic(err)
		}
		fmt.Println(num)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
