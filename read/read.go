/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the path to the file:")
	var fileName string
	_, scanErr := fmt.Scan(&fileName)
	if scanErr != nil {
		panic(scanErr)
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	results := make([]fullName, 0, 3)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		if len(split) != 2 {
			panic("Line should be split to 2 components only")
		}
		fullName := fullName{fname: split[0], lname: split[1]}
		results = append(results, fullName)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result.fname, result.lname)
	}
}

type fullName struct {
	fname string
	lname string
}
