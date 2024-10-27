package main

import (
	"fmt"
	f "functions/Functions"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		return
	}
	args := strings.Split(os.Args[len(os.Args)-1], " ")
	expressions := strings.Split(os.Args[len(os.Args)-2], "\\n")
	stack, err := f.ParseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if os.Args[1] == "checker" {
		if sort.IntsAreSorted(stack) && expressions == nil {
			fmt.Println("OK")
		} else if expressions != nil {
			stack = f.Swaper(stack, expressions)
			if sort.IntsAreSorted(stack) {
				fmt.Println("OK")
			} else {
				fmt.Println("KO")
			}
		} else {
			fmt.Println("KO")
		}
	}
}
