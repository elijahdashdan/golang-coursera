package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	fmt.Println("Please enter a name:")
	fmt.Scanln(&name)
	fmt.Println("Please enter an address:")
	fmt.Scanln(&address)

	info := make(map[string]string)
	info["name"] = name
	info["address"] = address

	jsonInfo, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonInfo))
}
