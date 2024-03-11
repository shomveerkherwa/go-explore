package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("darwin " + os)
	case "linux":
		fmt.Println("linux " + os)
	default:
		fmt.Println("may be windows " + os)
	}
}
