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
	var expenses []int64
	for i := 0; i < len(strExpenses); i++ {
		n, _ := strconv.ParseInt(strExpenses[i], 0, 64)
		expenses = append(expenses, n)
	}
	return expenses
}

// Part1 solution to Day1 part 1
func Part1() int64 {
	content := readFile("inputs/day1.txt")
	expensesStr := strings.Split(string(content), "\n")
	expenses := convertToInteger(expensesStr)

	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			valueA := expenses[i]
			valueB := expenses[j]
			if valueA+valueB == 2020 {
				fmt.Println(valueA, valueB)
				return valueA * valueB
			}
		}
	}

	return -1

}
