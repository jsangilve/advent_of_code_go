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

func Part2WithChannels() int {
	slopes := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	results_chan := make(chan int, len(slopes))
	defer func() {
		fmt.Println("Closing channel", results_chan)
		close(results_chan)
	}()

	for _, slope := range slopes {
		go func(slope [2]int) {
			results_chan <- countTrees(slopesMap, slope[0], slope[1])
		}(slope)
	}
	total := 1
	for i := 0; i < len(slopes); i++ {
		value := <-results_chan
		total *= value
	}
	return total
}

// func Part2WithChannelsV2() int {
// 	slopes := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
// 	results_chan := make(chan int, len(slopes))
// 	defer func() {
// 		fmt.Println("Closing channel", results_chan)
// 		// close(results_chan)
// 	}()

// 	for _, slope := range slopes {
// 		go func(slope [2]int) {
// 			results_chan <- countTrees(slopesMap, slope[0], slope[1])
// 		}(slope)
// 	}

// 	total := 1
// 	for {
// 		value, ok := <-results_chan
// 		fmt.Println("Value", value, ok, "numOfGoroutines", runtime.NumGoroutine())
// 		if !ok {
// 			fmt.Println("Channel", ok)
// 			break
// 		}
// 		total *= value
// 	}

// 	return total

// }
