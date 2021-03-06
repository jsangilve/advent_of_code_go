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
		Value1:   parseIntValue(matched[0]),
		Value2:   parseIntValue(matched[1]),
		Char:     matched[2],
		Password: matched[3],
	}
}

type passwordPolicy struct {
	Value1   int64
	Value2   int64
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
		// Value1 is Min and Value2 is Max
		if charOccurrences >= policy.Value1 && charOccurrences <= policy.Value2 {
			validPasswords = append(validPasswords, policy.Password)
		}
	}

	return len(validPasswords)
}

func safeCharAtPos(password string, index int) string {
	if len(password) > index {
		return string(password[index])
	}

	return "invalid"
}

// Par2 day2 solution to part 1
func Part2() int {
	content := utils.ReadFile("inputs/day2.txt")
	rows := strings.Split(string(content), "\n")
	var validPasswords []string
	for _, row := range rows {
		policy := parseRow(row)
		pos1 := safeCharAtPos(policy.Password, int(policy.Value1)-1)
		pos2 := safeCharAtPos(policy.Password, int(policy.Value2)-1)
		hasExactlyOne := (pos1 == policy.Char && pos2 != policy.Char) || (pos2 == policy.Char && pos1 != policy.Char)
		if hasExactlyOne {
			validPasswords = append(validPasswords, policy.Password)
		}

	}

	return len(validPasswords)

}
