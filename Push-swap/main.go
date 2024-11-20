package main

import (
	"fmt"
	"os"
	"strings"

	f "functions/Functions"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := strings.Split(os.Args[1], " ")
	stackA := f.ParseArgs(args)
	stackB, instructions := []int{}, []string{}
	f.PushSwap(&stackA, &stackB, &instructions)
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
	// fmt.Println(stackA, stackB)
}
