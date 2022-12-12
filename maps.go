package main

import "fmt"



func main(){
	var m=map[string]string{
		"a":"shakoor",
		"b":"machu",
	};
	fmt.Println(m["a"])
	m["c"]="riyad"

	
	fmt.Println(m)


}