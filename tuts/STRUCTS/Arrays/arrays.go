package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11}
	fmt.Println(primes)

	// a slice is an array without fixed size
	var slice = primes[0:3] // include first element , exclude last mentioned index
	fmt.Println(slice)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	arr1 := names[0:2]
	arr2 := names[1:3]
	fmt.Println(arr1, arr2)

	arr2[0] = "XXX"
	fmt.Println(arr1, arr2)
	fmt.Println(names)

	// slic defaults

}
