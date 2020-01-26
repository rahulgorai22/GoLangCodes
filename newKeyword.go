package main

import "fmt"
func main(){
	v:=new(int)
	*v++
	fmt.Println(*v)
}
