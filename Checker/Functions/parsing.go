package functions

import (
	"fmt"
	"os"
	"strconv"
)

func ParseArgs(args []string) []int {
	stackA := []int{}
	numbers := map[int]bool{}
	for _, arg := range args {
		if arg != "" {
			value, err := strconv.Atoi(arg)
			if err != nil || numbers[value] {
				fmt.Println("Error")
				os.Exit(0)
			}
			numbers[value] = true
			stackA = append(stackA, value)
		}
	}
	return stackA
}
