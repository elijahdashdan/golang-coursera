package main

import "fmt"

func main() {
	fmt.Println("Enter any floating number!")
	var n float64
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	i := int(n)
	fmt.Printf("The integer part is %d\n", i)
}
