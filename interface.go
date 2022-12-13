package main

import "fmt"

//interface declaration
type geometry interface{
	area() float64
	perim() float64

}

type sqr struct{
	length float64
}
type crl struct{
	radius float64
}
func (s sqr) area()float64{
	return s.length*s.length
}
func (s sqr) perim()float64{
	return s.length*4
}
func (c crl) area()float64{
	return c.radius*c.radius*3.14
}
func (c crl) perim()float64{
	return 2*c.radius*3.14
}
func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}
func main(){
	s := sqr{length: 4}
    c := crl{radius: 5}
	measure(s)
    measure(c)
}