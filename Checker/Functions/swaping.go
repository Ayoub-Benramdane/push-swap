package functions

func Swaper(stackA []int, instructions []string) []int {
	stackB := []int{}
	for _, c := range instructions {
		if c == "pa" {
			stackA = append([]int{stackB[0]}, stackA...)
			stackB = stackB[1:]
		} else if c == "pb" {
			stackB = append([]int{stackA[0]}, stackB...)
			stackA = stackA[1:]
		} else if c == "sa" {
			stackA[0], stackA[1] = stackA[1], stackA[0]
		} else if c == "sb" {
			stackB[0], stackB[1] = stackB[1], stackB[0]
		} else if c == "ss" {
			stackA[0], stackA[1] = stackA[1], stackA[0]
			stackB[0], stackB[1] = stackB[1], stackB[0]
		} else if c == "ra" {
			stackA = append(stackA[1:], stackA[:1]...)
		} else if c == "rb" {
			stackB = append(stackB[1:], stackB[:1]...)
		} else if c == "rr" {
			stackA = append(stackA[1:], stackA[:1]...)
			stackB = append(stackB[1:], stackB[:1]...)
		} else if c == "rra" {
			stackA = append(stackA[len(stackA)-1:], stackA[:len(stackA)-1]...)
		} else if c == "rrb" {
			stackB = append(stackB[len(stackB)-1:], stackB[:len(stackB)-1]...)
		} else if c == "rrr" {
			stackA = append(stackA[len(stackA)-1:], stackA[:len(stackA)-1]...)
			stackB = append(stackB[len(stackB)-1:], stackB[:len(stackB)-1]...)
		}
	}
	return stackA
}
