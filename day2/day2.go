package day2

import (
	"adventofcode/utils"
	"regexp"
	"strconv"
	"strings"
)

func parseIntValue(value string) int64 {
	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1
	}
	return n
}

func parseRow(row string) *passwordPolicy {
	matched := regexp.MustCompile(`[-\s:]+`).Split(row, -1)
	return &passwordPolicy{
		Min:      parseIntValue(matched[0]),
		Max:      parseIntValue(matched[1]),
		Char:     matched[2],
		Password: matched[3],
	}
}

type passwordPolicy struct {
	Min      int64
	Max      int64
	Char     string
	Password string
}

func countCharOcurrences(word string) map[string]int64 {
	var occurrences = make(map[string]int64)
	for _, rune := range word {
		char := string(rune)
		occurrences[char]++
	}
	return occurrences
}

// Part1 day2 solution to part 1
func Part1() int {
	content := utils.ReadFile("inputs/day2.txt")
	rows := strings.Split(string(content), "\n")
	var validPasswords []string
	for _, row := range rows {
		policy := parseRow(row)
		occurrences := countCharOcurrences(policy.Password)
		charOccurrences := occurrences[policy.Char]
		if charOccurrences >= policy.Min && charOccurrences <= policy.Max {
			validPasswords = append(validPasswords, policy.Password)
		}
	}

	return len(validPasswords)
}
