package day1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	return content
}

func convertToInteger(strExpenses []string) []int64 {
	expenses := make([]int64, len(strExpenses))
	for i, s := range strExpenses {
		n, _ := strconv.ParseInt(s, 0, 64)
		expenses[i] = n
	}
	return expenses
}

// Part1 solution to Day1 part 1
func Part1() int64 {
	content := readFile("inputs/day1.txt")
	expensesStr := strings.Split(string(content), "\n")
	expenses := convertToInteger(expensesStr)

	for i, valueA := range expenses {
		for j := i + 1; j < len(expenses); j++ {
			valueB := expenses[j]
			if valueA+valueB == 2020 {
				fmt.Println(valueA, valueB)
				return valueA * valueB
			}
		}
	}

	return -1

}
