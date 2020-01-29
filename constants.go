package main

import (
	"fmt"
	"math"

	//"math"
)
const s string="Rahul Gorai"
var p=fmt.Println
func main(){
	fmt.Println("Hello")
	const x=4000
	fmt.Println(x)
	const d=3e20/x
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(x))

	// Go follows strict typing, one type can't be assigned to other.
	//var i int32
	//var j int8
	//i=j // Not allowed
	i := 5
	j := i
	fmt.Println(j)



}
