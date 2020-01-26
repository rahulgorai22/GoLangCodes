package main

import "fmt"
func main(){
	//fmt.Printf("%*d",10,20)
	for i:=5;i>0;i--{
		fmt.Printf("%*d",i,10)
		fmt.Printf("\n")
	}
}
