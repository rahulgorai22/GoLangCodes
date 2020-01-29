package main

import (
	"fmt"
)

var (
	a int = b+1
	b=2
	c=getValue()
)
func getValue() int {
	return a+1
}
func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}