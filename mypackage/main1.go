package mypackage

import (
	"fmt"
)

var num int64 = 10

type cc func(int,int) int

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	//sayHello()
	//intSum(2,3)
	//intSum1(2,3)
	//ret := intSum(1, 2)
	//sum,sub := calc1( 2, 1)
	//fmt.Println(sum,sub)
	//var c cc
	//c = add
	//fmt.Println(c(1,2))

	//add := func(x,y int) {
		//fmt.Println(x + y)
	//}

	//add(10,20)

	//func(x, y int){
		//fmt.Println(x * y)
	//}(10, 20)

	var f = adder()
	fmt.Println(f(10))
	fmt.Println(f(10))

	var f1 = adder()
	fmt.Println(f1(10))
	fmt.Println(f1(10))
	fmt.Println(f(10))
}

func intSum(x, y int) int {
	return x + y
}

func sayHello()  {
	fmt.Println(num)
}

func intSum1(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _,v := range x {
		sum += v
	}
	return sum
}

func calc(x, y int) (int,int) {
	sum := x + y
	sub := x - y

	return sum,sub
}

func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y

	return
}

func add(x,y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}