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
	stackA, err := f.ParseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	} else if len(stackA) < 2 {
		fmt.Println("Error")
		return
	}
	stackB, instructions := []int{}, []string{}
	tot := len(stackA)
	instructions, _, _ = f.PushSwap(stackA, stackB, instructions, tot)
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}
