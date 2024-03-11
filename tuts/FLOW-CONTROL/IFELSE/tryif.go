package main

import (
	"fmt"
	"math"
)

func pow(x, y, limit float64) float64 {
	if result := math.Pow(x, y); result < limit {
		return result
	} else {
		fmt.Printf("%v is above the limit %v", result, limit)
	}
	return -1
}

func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))
}
