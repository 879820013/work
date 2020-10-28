package df

import "fmt"

func NewF()  {
	/*fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")*/
	a := 1
	fmt.Println(a)

	defer fmt.Println(a)

	a += 1
}