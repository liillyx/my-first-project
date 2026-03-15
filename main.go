package main

import "fmt"

func main() {
	number := 7
	result := square(number)
	fmt.Println(result)
}
func square(n int) int {
	return n * n
}
func cube(n int) int {
	return n * n * n
}
