package main

import (
	"adventofcode/day1"
	"adventofcode/day2"
	"adventofcode/day3"
	"fmt"
)

func main() {
	fmt.Println("Day1 part 1", day1.Part1())
	fmt.Println("Day1 part 2", day1.Part2())

	fmt.Println("Day2 part1", day2.Part1())
	fmt.Println("Day2 part2", day2.Part2())

	fmt.Println("Day3 part1", day3.Part1())
	fmt.Println("Day3 part2", day3.Part2())
	fmt.Println("Day3 part2 with channels", day3.Part2WithChannels())
	fmt.Println("Day3 part2 with channels v2", day3.Part2WithChannelsV2())
	fmt.Println("Day3 part2 with channels v3", day3.Part2WithChannelsV3())

}
