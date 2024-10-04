package main

import (
	"fmt"
	f "functions/Functions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := strings.Split(os.Args[1], " ")
	stackA, err := f.ParseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, instruction := range f.PushSwap(stackA) {
		fmt.Println(instruction)
	}
}
