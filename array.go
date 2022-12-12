package main

import (
	"fmt"
)

func main(){
	fmt.Println("welcome")
	var array [5]int;
	array[0]=1
	array[1]=2
	array[2]=3

	fmt.Println(array[1])

	var array1 [3]int=[3]int{1,2,3}
	fmt.Println(array1[2])

	k := [...]int{10, 20, 30}
	fmt.Println(len(k))


	//COPY AN ARRAY

	array2:=array
	fmt.Printf("%v",array2)

	//COPY BY REFERENCE

	z:=&array
	fmt.Printf("z: %v\n", *z)
	
}