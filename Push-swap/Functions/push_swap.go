package functions

import "sort"

func PushSwap(stackA, stackB []int, instructions []string) ([]string, []int, []int) {
	if sort.IntsAreSorted(stackA) {
		return instructions, stackA, stackB
	}
	if len(stackA) > 3 {
		min := stackA[0]
		count := 0
		for i := 1; i < len(stackA); i++ {
			if stackA[i] < min {
				count = i
				min = stackA[i]
			}
		}
		stackB = append(stackB, stackA[count])
		stackA = append(stackA[:count], stackA[count+1:]...)
		pos := count
		for count > 0 {
			count--
			if pos > len(stackA)/2 {
				instructions = append(instructions, "rra")
			} else {
				instructions = append(instructions, "ra")
			}
		}
		instructions = append(instructions, "pb")
		if len(stackA) > 3 {
			instructions, stackA, stackB = PushSwap(stackA, stackB, instructions)
		}
	}
	if len(stackA) <= 3 {
		if stackA[0] < stackA[1] && stackA[0] < stackA[2] && stackA[1] > stackA[2] {
			stackA[2], stackA[1] = stackA[1], stackA[2]
			instructions = append(instructions, "rra", "sa")
		} else if stackA[0] > stackA[1] && stackA[0] < stackA[2] && stackA[1] < stackA[2] {
			stackA[0], stackA[1] = stackA[1], stackA[0]
			instructions = append(instructions, "sa")
		} else if stackA[0] < stackA[1] && stackA[1] > stackA[2] {
			stackA[0], stackA[2] = stackA[2], stackA[0]
			instructions = append(instructions, "rra")
		} else if stackA[0] > stackA[1] && stackA[1] < stackA[2] {
			stackA[0], stackA[2] = stackA[2], stackA[0]
			stackA[1], stackA[0] = stackA[0], stackA[1]
			instructions = append(instructions, "ra")
		} else if stackA[0] > stackA[1] && stackA[1] > stackA[2] {
			stackA[0], stackA[2] = stackA[2], stackA[0]
			instructions = append(instructions, "sa", "rra")
		}
		for i := len(stackB) - 1; i >= 0; i-- {
			stackA = append(stackA, stackB[i])
			instructions = append(instructions, "pa")
		}
	}
	return instructions, stackA, stackB
}
