package main

import (
	"fmt"
	"math"
)

func main() {
	var discountPercent float64 = 20.0
	var productPrice float64 = 100.0

	newPrice := productPrice - (productPrice * (discountPercent / 100))
	finalPrice := math.Floor(newPrice*100) / 100
	fmt.Println(finalPrice)

}
