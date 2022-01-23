package main

import (
	"fmt"
	"os"
)

func main() {
	var num float32

	fmt.Println("Enter the floating point number")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input is not parsable as float")
		os.Exit(-1)
	}
	fmt.Println(int(num))

}
