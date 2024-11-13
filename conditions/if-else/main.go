package main

import "fmt"

func main() {
	fmt.Println("Enter your number")
	var x int
	fmt.Scanln(&x)
	if x > 10 {
		fmt.Println("greater than 10")
	} else if x==10{
		fmt.Println("equals to 10")
	} else {
		fmt.Println("lower than 10")
	}

	
}