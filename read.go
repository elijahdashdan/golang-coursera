package main

import (
	"bufio"
	"fmt"
	"os"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Println("Please enter the path of the text file:")
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var names []name
	for scanner.Scan() {
		fields := name{}
		fmt.Sscanf(scanner.Text(), "%s %s", &fields.fname, &fields.lname)
		names = append(names, fields)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}
}
