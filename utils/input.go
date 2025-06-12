package utils

import "fmt"

func Input(prompt string) int {
	fmt.Print(prompt)
	var input int
	fmt.Scan(&input)
	return input
}

func InputTemp(prompt string) float64 {
	fmt.Print(prompt)
	var input float64
	fmt.Scan(&input)
	return input
}
