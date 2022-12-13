package main

import "fmt"

func main(){
	i:=20
	p:=&i
	fmt.Println(p)

	// Set i using the pointer
	*p = 21
	fmt.Println(i)
}