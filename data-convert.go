package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	var a string="3.22"
	// Convert String to Float
	f, _ := strconv.ParseFloat(a,8)
	var s float64 =f+22.22
	fmt.Println(s)

	//conver string to boolean

	b:="true"
	
	t,_:=strconv.ParseBool(b)
	fmt.Println(t);


	// Converting float to int
	var f32 float32 = 3.1415926535
	fmt.Println(reflect.TypeOf(f32))
	i32 := int32(f32)
	fmt.Println(reflect.TypeOf(i32))

	// Float to String
	var flo float64 = 3.1415926535
	var strFlo string = strconv.FormatFloat(flo, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(strFlo))
	fmt.Println(strFlo)
}