package df

import "fmt"

/**
* "A",10,20,30
* "AA",10,30,40
* "B",10,20,30
* "BB",10,30,40
**/
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func Neww()  {
	x := 1
	y := 2
	defer calc("AA",x,calc("A",x,y))
	x = 10
	defer calc("BB",x,calc("B",x,y))
	y = 20
}
