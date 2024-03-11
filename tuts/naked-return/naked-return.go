package main

import "fmt"

// written short functions to improve readability
func getValues(x, y string) (z string) {
	z = x + y
	return
}

func main() {
	fmt.Println(getValues("jan", "go"))
}
