package main

import "fmt"

func swap(x,y string) (string,string) {
	return y,x;
}

func main(){
	a,b := swap("som","veer");
	fmt.Println(a,b);
	fmt.Println(b,a);
}