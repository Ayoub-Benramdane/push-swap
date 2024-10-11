package main

import (
	"fmt"
	f "functions/Functions"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := strings.Split(os.Args[2], " ")
	stack, err := f.ParseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if os.Args[1] == "checker" {
		if sort.IntsAreSorted(stack) {
			fmt.Println("OK")
		} else {
			fmt.Println("KO")
		}
	}
}
