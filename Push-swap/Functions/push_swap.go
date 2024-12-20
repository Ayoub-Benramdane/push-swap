package functions

import (
	"sort"
)

func PushSwap(stackA, stackB *[]int, instructions *[]string) {
	if sort.IntsAreSorted((*stackA)) {
		return
	} else if len(*stackA) > 3 {
		count, pos, checker := minNumber(stackA)
		moveToTop(stackA, instructions, count, pos, checker)
		if !sort.IntsAreSorted(*stackA) {
			*stackB = append([]int{(*stackA)[0]}, *stackB...)
			(*stackA) = append([]int{}, (*stackA)[1:]...)
			*instructions = append(*instructions, "pb")
			PushSwap(stackA, stackB, instructions)
		}
	} else if len(*stackA) < 3 && (*stackA)[1] < (*stackA)[0] {
		(*stackA)[0], (*stackA)[1] = (*stackA)[1], (*stackA)[0]
		*instructions = append(*instructions, "sa")
	} else if len(*stackA) == 3 {
		smallStack(stackA, instructions)
	}
	for i := 0; len(*stackB) != 0; i++ {
		(*stackA) = append([]int{(*stackB)[i]}, (*stackA)...)
		(*stackB) = (*stackB)[i+1:]
		i--
		*instructions = append(*instructions, "pa")
	}
}

func minNumber(stackA *[]int) (int, int, int) {
	min, checker, count := (*stackA)[0], (*stackA)[0], 0
	for i := 1; i < len(*stackA); i++ {
		if (*stackA)[i] < min {
			checker = min
			count = i
			min = (*stackA)[i]
		}
		if (*stackA)[i] > min && (*stackA)[i] < checker {
			checker = (*stackA)[i]
		}
	}
	return count, count, checker
}

func smallStack(stackA *[]int, instructions *[]string) {
	if (*stackA)[0] < (*stackA)[1] && (*stackA)[0] < (*stackA)[2] && (*stackA)[1] > (*stackA)[2] {
		(*stackA)[2], (*stackA)[1] = (*stackA)[1], (*stackA)[2]
		*instructions = append(*instructions, "rra", "sa")
	} else if (*stackA)[0] > (*stackA)[1] && (*stackA)[0] < (*stackA)[2] && (*stackA)[1] < (*stackA)[2] {
		(*stackA)[0], (*stackA)[1] = (*stackA)[1], (*stackA)[0]
		*instructions = append(*instructions, "sa")
	} else if (*stackA)[0] < (*stackA)[1] && (*stackA)[1] > (*stackA)[2] {
		(*stackA)[0], (*stackA)[2] = (*stackA)[2], (*stackA)[0]
		*instructions = append(*instructions, "rra")
	} else if (*stackA)[0] > (*stackA)[1] && (*stackA)[1] < (*stackA)[2] {
		(*stackA)[0], (*stackA)[2] = (*stackA)[2], (*stackA)[0]
		(*stackA)[1], (*stackA)[0] = (*stackA)[0], (*stackA)[1]
		*instructions = append(*instructions, "ra")
	} else if (*stackA)[0] > (*stackA)[1] && (*stackA)[1] > (*stackA)[2] {
		(*stackA)[0], (*stackA)[2] = (*stackA)[2], (*stackA)[0]
		*instructions = append(*instructions, "sa", "rra")
	}
}

func moveToTop(stackA *[]int, instructions *[]string, count, pos, checker int) {
	for count > 0 && count < len(*stackA) {
		if pos > len(*stackA)/2 {
			*instructions = append(*instructions, "rra")
			(*stackA) = append((*stackA)[len(*stackA)-1:], (*stackA)[:len(*stackA)-1]...)
			count++
		} else if (*stackA)[0] == checker && count == 1 {
			*instructions = append(*instructions, "sa")
			(*stackA)[0], (*stackA)[1] = (*stackA)[1], (*stackA)[0]
			count--
		} else {
			*instructions = append(*instructions, "ra")
			(*stackA) = append((*stackA)[1:], (*stackA)[0])
			count--
		}
	}
}
