package main

import "fmt"

func main() {
	result := sum(2, 5)
	fmt.Println(result)
}

func sum(n1, n2 int) int {
	result := n1 + n2
	return result
}
