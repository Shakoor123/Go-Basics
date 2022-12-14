package main 

import (
	"fmt"
	"errors"
	"math"

)
func Sqrt(x float64)(float64 , error){
	if x < 0 {
		return 0, errors.New("Math: Not possible to calculate the square root of negative number")
	}else{
		z := 1.0
		min_delta := 0.00000000001
		delta := 1.0
		i := 0
		for ; math.Abs(delta) > min_delta; i++ {
			delta = (z*z - x) / (2 * z)
			z = z - delta
		}
		fmt.Println("iterations: ", i)
		return z, nil
	}
}


func main(){
	//creating an error
	err :=errors.New("first error")
	fmt.Println("Error : ", err)

	//get an error as return value

	f,err:=Sqrt(2)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(f)
	}
}