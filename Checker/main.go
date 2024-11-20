package main

import (
	"fmt"
	f "functions/Functions"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := strings.Split(os.Args[len(os.Args)-1], " ")
	var expressions string
	stack := f.ParseArgs(args)
	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}
		expressions += input +"\\n"
	}
	sliceExpressions := strings.Split(expressions, "\\n")
	stack = f.Swaper(stack, sliceExpressions)
	if sort.IntsAreSorted(stack) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
			
	}
}
