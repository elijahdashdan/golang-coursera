package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)
	for {
		var input string
		fmt.Print("Enter an integer (or 'X' to exit): ")
		fmt.Scanln(&input)

		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		slice = append(slice, num)
		sort.Ints(slice)
		fmt.Println(slice)
	}
}
