package main

import (
	"fmt"
)

func main(){
	// var a int=22
	// if(a%2==0){
	// 	fmt.Println("Even")
	// }else{
	// 	fmt.Println("Odd")
	// }


	// switch(a){
	// case 11:fmt.Println(a)
	// break;
	// case 22:fmt.Println(a)
	// break;
	// default:fmt.Println("error")
	// }

	// var i int;
	// for i=0;i<=10;i++ {
	// 	fmt.Println(i)
	// }
	
	
	// for i<=10 {
	// 	fmt.Println(i)
	// 	i++;
	// }

	num:=5;
	for{
		fmt.Println(num)
		if num==10{
			break
		}
		num++
	}
}