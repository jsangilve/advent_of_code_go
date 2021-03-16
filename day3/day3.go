package day3

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func parseMap() [][]string {
	content := utils.ReadFile("inputs/day3.txt")
	rows := strings.Split(string(content), "\n")
	slopesMap := make([][]string, len(rows))
	for i, strRow := range rows {
		slopesMap[i] = strings.Split(strRow, "")
	}
	return slopesMap

}

func countTrees(slopesMap [][]string, right int, down int) int {
	rowLength := len(slopesMap[0])
	x, y := 0, 0
	numTrees := 0
	for y < len(slopesMap) {
		if slopesMap[y][x%rowLength] == "#" {
			numTrees++
		}
		x += right
		y += down
	}
	return numTrees

}

var slopesMap = parseMap()

func Part1() int {
	return countTrees(slopesMap, 3, 1)
}

func Part2() int {
	slope1 := countTrees(slopesMap, 1, 1)
	slope2 := countTrees(slopesMap, 3, 1)
	slope3 := countTrees(slopesMap, 5, 1)
	slope4 := countTrees(slopesMap, 7, 1)
	slope5 := countTrees(slopesMap, 1, 2)
	fmt.Println(slope1, slope2, slope3, slope4, slope5)
	return slope1 * slope2 * slope3 * slope4 * slope5
}
