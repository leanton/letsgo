/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	res := make(map[string]string)
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name")
	res["name"] = readLine(in)
	fmt.Println("Enter the address")
	res["address"] = readLine(in)

	marshal, jsonErr := json.Marshal(res)
	if jsonErr != nil {
		panic(jsonErr)
	}
	fmt.Println(string(marshal))
}

func readLine(in *bufio.Reader) string {
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	return line
}
