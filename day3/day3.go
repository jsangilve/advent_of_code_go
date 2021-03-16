package day3

import (
	"adventofcode/utils"
	"strings"
)

func parseMap() [][]string {
	content := utils.ReadFile("inputs/day3.txt")
	rows := strings.Split(string(content), "\n")
	slopesMap := make([][]string, len(rows))
	for i, strRow := range rows {
		slopesMap[i] = strings.Split(strRow, "")
		// fmt.Println(slopesMap[i])
	}
	return slopesMap

}

func Part1() int {
	slopesMap := parseMap()
	rowLength := len(slopesMap[0])
	x, y := 0, 0
	numTrees := 0
	for y < len(slopesMap) {
		if slopesMap[y][x%rowLength] == "#" {
			numTrees++
		}
		x += 3
		y++
	}
	return numTrees
}
