package main

import (
	"advent-of-code-2022/util"
	"fmt"
	"strconv"
	"strings"
)

var (
	biggerCalorie int64 = 0
)

func main() {
	input := util.ReadInput("/dev/work/advent-of-code-2022/day-1/input.txt")
	elves := strings.Split(input, "\r\n\r")
	for _, elf := range elves {
		calories := strings.Split(elf, "\r")
		var elfCalories int64 = 0
		for _, calorie := range calories {
			calorie = strings.ReplaceAll(calorie, "\n", "")
			calorieInt, _ := strconv.ParseInt(calorie, 10, 0)
			elfCalories += calorieInt
		}
		if biggerCalorie < elfCalories {
			biggerCalorie = elfCalories
		}
	}
	fmt.Printf("The elf with more calories of food have %d", biggerCalorie)
}
