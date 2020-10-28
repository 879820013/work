package df

import "fmt"

func funA()  {
	fmt.Println("func A")
}

func funB()  {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funC()  {
	fmt.Println("func C")
}

func NewError()  {
	funA()
	funB()
	funC()
}
