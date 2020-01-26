package main

import (
	"fmt"
)

type point struct {
	x,y int
}
func main(){
	p1:=point{1,2}
	fmt.Printf("The Y cordinate is %[2]d and the X coordiante is %[1]d",p1.x,p1.y)
	fmt.Printf("%s\n","rahul")
	fmt.Printf("%x\n","A")
}