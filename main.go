package main

import (
	"fmt"
	"math"
)

func main() {
	var random float64 = 44.4
	isEven := random == math.Trunc(random) && int(random)%2 == 0

	fmt.Printf("Исходное число: %.1f\n", random)
	fmt.Printf("Исходное число, увеличенное на 10%%: %.5f\n", random+(random*0.1))
	fmt.Printf("Исходное число является четным: %t\n", isEven)
	if random < 10 {
		fmt.Println("Предпоследняя цифра целой части исходного числа: 0")
	} else if 9 < random && random < 100 {
		fmt.Printf("Предпоследняя цифра целой части исходного числа: %d\n", int(random)/10)
	} else if random > 99 {
		fmt.Printf("Предпоследняя цифра целой части исходного числа: %d\n", int(random)/100)
	}
}
