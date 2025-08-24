package main

import "fmt"

func main() {
	var num int
	for i := 0; num < 100; i++ {
		fmt.Scan(&num)
		if num > 100 {
			break
		} else if num < 10 {
			continue
		} else {
			fmt.Println(num)
		}
	}
}
