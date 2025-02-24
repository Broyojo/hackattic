package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		weekdayNum := num % 7
		if weekdayNum < 0 {
			weekdayNum += 7
		}
		epoch := time.Unix(0, 0)
		weekday := (epoch.Weekday() + time.Weekday(weekdayNum)) % 7
		fmt.Println(weekday)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
