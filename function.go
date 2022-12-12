package main

import (
	"fmt"
)

// func main(){
// 	var a int=10;
// 	var b int=20;
// 	c:=hey(a,b)
// 	fmt.Println(c)
// }
// func hey(a int,b int)int {
// 	var c int=a+b;
// 	return c;
// }

func main(){
	var a int=10;
	var b int=20;
	hey(a,b)
}
func hey(a int,b int) {
	fmt.Println(a+b)
}