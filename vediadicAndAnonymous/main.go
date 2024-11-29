package main

import "fmt"

func main() {
	fmt.Println("sum=", add(5, 6, 7))

	// jb function ke andr function bnate h to wo anonyous function kehlaya jaata h
	// syntax

	// func() {
	//
	// }()

	// now implementation of anonymous function

	func(a int) int {
		fmt.Println("anonymous function running")
		return 0

	}(2)
}

func add(a ...int) int {
	// here ... shows that a can hold different values and to use that we need to use range keyword
	// mtlb b ke andr ki value ko ek ek krke nikaal rhe h or sum me add kr rhe h
	sum := 0
	for _, val := range a {
		sum += val
	}
	return sum

}
