package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		line := scanner.Text()

		var decoded []int64
		for _, part := range strings.Split(line, " ") {
			switch {
			case len(part) > 2 && part[0] == '0' && part[1] == 'x':
				value, err := strconv.ParseInt(part[2:], 16, 64)
				if err != nil {
					panic(err)
				}
				decoded = append(decoded, value)
			case len(part) > 2 && part[0] == '0' && part[1] == 'b':
				value, err := strconv.ParseInt(part[2:], 2, 64)
				if err != nil {
					panic(err)
				}
				decoded = append(decoded, value)
			case len(part) > 2 && part[0] == '0' && part[1] == 'o':
				value, err := strconv.ParseInt(part[2:], 8, 64)
				if err != nil {
					panic(err)
				}
				decoded = append(decoded, value)
			case isInt(part):
				value, err := strconv.ParseInt(part, 10, 64)
				if err != nil {
					panic(err)
				}
				decoded = append(decoded, value)
			case len(part) == 1:
				// ascii digit
				value := int64(part[0])
				decoded = append(decoded, value)
			}
		}

		var sum int64
		for _, d := range decoded {
			sum += d
		}
		fmt.Println(sum)
	}
}
