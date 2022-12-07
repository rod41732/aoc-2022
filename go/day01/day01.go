package main

import (
	"fmt"
	"sort"
	"strconv"
  "aoc/utils"
  "strings"
)


func parse(file string) []int {
  var calories = make([]int, 0)
  for _, chunk := range strings.Split(file, "\n\n") {
    var sum = 0
    for _, line := range strings.Split(chunk, "\n") {
      if (line == "") {continue}
      sum += utils.Unwrap(strconv.Atoi(line))
    }
    calories = append(calories, sum)
  }
  sort.Ints(calories)
  utils.Reverse(calories)
  return calories
}

func part1(sortedDescCalories []int) int {
  return sortedDescCalories[0]
}

func part2(sortedDescCalories []int) int {
  return sortedDescCalories[0]+sortedDescCalories[1] +sortedDescCalories[2]
}

func main() {
  // var acc = e([]int, 0)
  var testCalories = parse(utils.ReadFile("day01/day01_test.txt"))
  utils.AssertEqual(part1(testCalories), 24000)
  utils.AssertEqual(part2(testCalories), 45000)

  var calories = parse(utils.ReadFile("day01/day01.txt"))
  fmt.Printf("Part 1: %d\n", part1(calories))
  fmt.Printf("Part 2: %d\n", part2(calories))

}
