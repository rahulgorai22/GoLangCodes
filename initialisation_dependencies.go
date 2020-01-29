package main

import "fmt"

// Independent variables are processed first , then dependent variable are processed outside the function.
// inside function it is evaluated as it is left to right and top to bottom
var (
	a int=b+1
	b=10
)
func main(){
	fmt.Println(a)
	fmt.Println(b)
}
