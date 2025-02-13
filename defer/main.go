// defer follows the stack rule
package main

import "fmt"

func main() {
	for i:=1 ; i<11 ; i++{
		fmt.Println(i)
	}
	fmt.Println("After using defer..........")

	for i:=1 ; i<11 ; i++{
		defer fmt.Println(i)
		// here in this case defer is used to print the number in reverse order
	}

	defer fmt.Println("hello")
	defer fmt.Println("world")
	// defer wale saare lines reverse order me chalte h

	// there is one usecase also agr code me error h to compiler bs defer ko compile karega


}
