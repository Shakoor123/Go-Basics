package main

import "fmt"

type user struct{
	name string
	age int
}

func main(){
	var student user;
	student.name="shakoor"
	student.age=21
	
	fmt.Println(student)
}