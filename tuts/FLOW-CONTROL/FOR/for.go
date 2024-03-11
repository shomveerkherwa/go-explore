package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + i
	}
	fmt.Printf("%v \n", sum)

	// init and post tatement are optional
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// using for as while loop
	count := 0
	for count <= 10 {
		fmt.Printf("hi")
		count++
	}

	// if statement also dont need braces but need semi-colon
	if count > 10 {
		fmt.Println(" \n Enough HI's")
	}

}
