package functions

import (
	"sort"
)

func PushSwap(stackA, stackB *[]int, instructions *[]string) {
	if len(*stackA) > 3 {
		if sort.IntsAreSorted((*stackA)[1:]) {
			if (*stackA)[0] > (*stackA)[len(*stackA)-1] {
				*instructions = append(*instructions, "ra")
				*stackA = append((*stackA)[1:len(*stackA)-1], (*stackA)[0], (*stackA)[len(*stackA)-1])
			} else {
				*instructions = append(*instructions, "rra", "pb", "ra", "pa", "ra")
				*stackA = append((*stackA)[1:], (*stackA)[:1]...)
			}
		} else if !sort.IntsAreSorted(*stackA) {
			count, pos := minNumber(stackA)
			moveToTop(stackA, instructions, count, pos)
			*stackB = append([]int{(*stackA)[0]}, *stackB...)
			(*stackA) = append([]int{}, (*stackA)[1:]...)
			*instructions = append(*instructions, "pb")
			if !sort.IntsAreSorted(*stackA) && len(*stackA) > 3 {
				PushSwap(stackA, stackB, instructions)
			}
		}
	}
	if len(*stackA) < 3 && (*stackA)[1] < (*stackA)[0] {
		(*stackA)[0], (*stackA)[1] = (*stackA)[1], (*stackA)[0]
		*instructions = append(*instructions, "sa")
	} else if len(*stackA) == 3 {
		smallStack(stackA, instructions)
	}
	for i := 0; i < len(*stackB); i++ {
		(*stackA) = append([]int{(*stackB)[i]}, (*stackA)...)
		*instructions = append(*instructions, "pa")
	}
}

func minNumber(stackA *[]int) (int, int) {
	min := (*stackA)[0]
	count := 0
	for i := 1; i < len(*stackA); i++ {
		if (*stackA)[i] < min {
			count = i
			min = (*stackA)[i]
		}
	}
	return count, count
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

func moveToTop(stackA *[]int, instructions *[]string, count, pos int) {
	if count == 1 && (*stackA)[0] < (*stackA)[2] && (*stackA)[0] > (*stackA)[len(*stackA)-1] {
		*instructions = append(*instructions, "sa")
		(*stackA)[0], (*stackA)[1] = (*stackA)[1], (*stackA)[0]
	} else {
		for count > 0 && count < len(*stackA) {
			if pos > len(*stackA)/2 {
				*instructions = append(*instructions, "rra")
				(*stackA) = append((*stackA)[len(*stackA)-1:], (*stackA)[:len(*stackA)-1]...)
				count++
			} else {
				*instructions = append(*instructions, "ra")
				(*stackA) = append((*stackA)[1:], (*stackA)[:1]...)
				count--
			}
		}
	}
}
