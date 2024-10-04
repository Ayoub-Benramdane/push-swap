package functions

import (
	"fmt"
	"strconv"
)

func ParseArgs(args []string) ([]int, error) {
	stackA := []int{}
	numbers := map[int]bool{}
	for _, arg := range args {
		if arg != "" {
			value, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			if numbers[value] {
				return nil, fmt.Errorf("duplicate number: %d", value)
			}
			numbers[value] = true
			stackA = append(stackA, value)
		}
	}
	return stackA, nil
}
