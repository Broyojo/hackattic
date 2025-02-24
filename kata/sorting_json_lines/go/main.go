package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Record map[string]map[string]any

func (r Record) getName() string {
	for k := range r {
		if k != "extra" {
			return k
		}
	}
	return ""
}

func (r Record) getBalance() int {
	if extra, ok := r["extra"]; ok {
		balance := extra["balance"]
		if value, ok := balance.(float64); ok {
			return int(value)
		}
	}

	for k, v := range r {
		if k != "extra" {
			if balance, ok := v["balance"]; ok {
				if value, ok := balance.(float64); ok {
					return int(value)
				}
			}
			break
		}
	}

	return -1
}

func formatInt(x int) string {
	str := strconv.Itoa(x)
	var withCommas string
	var digits int
	for i := len(str) - 1; i >= 0; i-- {
		if digits == 3 && i != len(str)-1 {
			digits = 0
			withCommas = "," + withCommas
		}
		withCommas = string(str[i]) + withCommas
		digits++
	}
	return withCommas
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var records []Record
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		var record Record
		if err := json.Unmarshal([]byte(scanner.Text()), &record); err != nil {
			panic(err)
		}

		records = append(records, record)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].getBalance() < records[j].getBalance()
	})

	for _, record := range records {
		fmt.Printf("%v: %v\n", record.getName(), formatInt(record.getBalance()))
	}
}
