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
	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}
		expressions += input +"\\n"
	}
	sliceExpressions := strings.Split(expressions, "\\n")
	stack, err := f.ParseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	stack = f.Swaper(stack, sliceExpressions)
	if sort.IntsAreSorted(stack) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
			
	}
}
