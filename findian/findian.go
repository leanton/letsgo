package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the string")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(-1)
	}

	lowerInput := strings.ToLower(input)
	if strings.HasPrefix(lowerInput, "i") && strings.HasSuffix(lowerInput, "n") && strings.Contains(lowerInput, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
