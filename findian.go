package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter any String.")
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	s := strings.ToLower(str)
	if s[0] == 'i' && s[len(s)-1] == 'n' && strings.Contains(s, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
