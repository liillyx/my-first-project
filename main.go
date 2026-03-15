package main

import "fmt"

func main() {
	number := 7
	fmt.Println(cube(number))
}
func square(n int) int {
	return n * n
}
func cube(n int) int {
	return n * n * n
}
