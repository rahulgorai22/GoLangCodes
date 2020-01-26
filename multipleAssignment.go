package main

import "fmt"
func main(){
	fmt.Println("Program to illustrate multiple assignment of variables in goLang")
	var a,b int=10,20
	a,b=b,a
	fmt.Println(a,b)
	// Takes multiple return values from function
	c,d:=sumProduct(1,2)
	fmt.Println(c,d)
}
func sumProduct (a,b int) (int,int){
	return a+b,a-b
}