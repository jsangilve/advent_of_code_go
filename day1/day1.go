package day1

import (
	"adventofcode/utils"
	"strconv"
	"strings"
)

func convertToInteger(strExpenses []string) []int64 {
	expenses := make([]int64, len(strExpenses))
	for i, s := range strExpenses {
		n, _ := strconv.ParseInt(s, 0, 64)
		expenses[i] = n
	}
	return expenses
}

func readExpenses() []int64 {
	content := utils.ReadFile("inputs/day1.txt")
	expensesStr := strings.Split(string(content), "\n")
	return convertToInteger(expensesStr)
}

// Part1 solution to Day1 part 1
func Part1() int64 {
	expenses := readExpenses()

	for i, valueA := range expenses {
		for j := i + 1; j < len(expenses); j++ {
			valueB := expenses[j]
			if valueA+valueB == 2020 {
				return valueA * valueB
			}
		}
	}

	return -1
}

// Part2 solution to Day 1 part 2
func Part2() int64 {
	expenses := readExpenses()
	numExpenses := len(expenses)

	for i := 0; i < numExpenses; i++ {
		valueA := expenses[i]
		for j := i + 1; j < numExpenses; j++ {
			valueB := expenses[j]
			for k := 0; k < numExpenses; k++ {
				valueC := expenses[k]
				if k != j && k != i && valueA+valueB+valueC == 2020 {
					return valueA * valueB * valueC
				}
			}
		}
	}
	return -1
}
