package main
import (
	"fmt"
)

type T struct {
	age int
}

func (t T) String() string{
	if t.age >50{
		return "senior"
	}
	return "junior"
}
func main()  {
	t:=T{30}
	fmt.Println(t)
	fmt.Printf("%s\n",t)

}
