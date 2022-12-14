package main

import "fmt"
func recoverFuction(){
	//recover from panic and print error message
	if recoveryMessage:=recover();recoveryMessage != nil{
		fmt.Println(recoveryMessage)
	}
	fmt.Println("End: recoveryFunction")
}
func executePanic(){
	//Defer func call

	defer recoverFuction()

	//Throw panic

	panic("panic")
	fmt.Println("End execute panic")
}
 

func main(){
	executePanic()
	fmt.Println("End main")
}