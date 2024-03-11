package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	// 2. access using dots
	v.Y = 4;
	fmt.Println(v)

	// 3. pointers can be used on structs too
	p := &v;
	p.X = 99;

	fmt.Println(v)

	// 4. populate only few fields
	v2 := Vertex {Y:99} // x will be 0, i.e it picks up default values
	fmt.Println(v2)

	v3 := Vertex{}
	fmt.Println(v3)
	
	
}