package main
import "fmt"
var p=fmt.Println
func main(){
	p("Program to illustrate operators and values in goLang")
	p("Go"+"Lang")
	p(1+2)
	p(1+2.3)
	p(2.5/2.2)
	p(true||true)
	p(true&&true)
	p(!true)
}
