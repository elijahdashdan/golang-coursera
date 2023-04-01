package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a string: ")
	//_, err := fmt.Scanln(&s)
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
	line = strings.Replace(line, "\n", "", -1)
	line = strings.ToLower(line)
	if strings.HasPrefix(line, "i") &&
		strings.HasSuffix(line, "n") &&
		strings.Contains(line, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
