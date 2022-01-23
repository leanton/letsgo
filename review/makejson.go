package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getUserInput(usr_inp_name, usr_inp string) string {
	fmt.Print(fmt.Sprintf("Enter the %s: ", usr_inp_name))
	scn, err := fmt.Scanf("%q\n", &usr_inp)
	if scn == 0 {
		fmt.Println("Empty input passed")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return usr_inp
}

func main() {

	var name, addr string
	name = getUserInput("name", name)
	addr = getUserInput("address", addr)

	user_map := map[string]string{
		"name":    name,
		"address": addr,
	}

	json_str, _ := json.Marshal(user_map)
	fmt.Println(string(json_str))

}
