package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	res := make([]int, 0, 3)
	for {
		var input string
		fmt.Print("Enter an integer to be added to the slice: ")
		_, scanErr := fmt.Scan(&input)
		if scanErr != nil {
			os.Exit(-1)
		}
		if strings.ToLower(input) == "x" {
			break
		}
		i, convErr := strconv.Atoi(input)
		if convErr != nil {
			os.Exit(-1)
		}
		res = append(res, i)
		sort.Ints(res)
		fmt.Println(res)
	}
}
