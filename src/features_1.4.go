package main

import "fmt"

func main() {
	fmt.Printf("Hello Go 1.4")

	x := []int{1, 2, 3}
	for range x {
		fmt.Printf(".")
	}
	fmt.Printf("all done!\n")
}
