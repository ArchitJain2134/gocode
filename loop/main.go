package main

import "fmt"

func main() {
	// normal for loop
	for i := 0; i < 4; i++ {
		fmt.Println("number is", i)
	}

	// infinite for loop
	counter := 0
	for {
		fmt.Println("Hello")
		counter++
		if counter == 3 {
			break
			// break statement helps in exiting the infinite loop
		}
	}

	// we also use range keyword in for loop
	// range keyword simplifies looping over element of a collections like slices,arrays,mapa and strings
	// using range in integer type array
	numbers := []int{135, 342, 456, 245, 645}
	for index, value := range numbers {
		// isme range keyword array ke har element ke paas jaa rha hai or uska index no or uski value bta rha hai
		fmt.Printf("index value is %d and the value is %d \n", index, value)
	}

	// using range in string type array
	data := "hello my name is archit jain"
	for index,value:=range data{
		fmt.Printf("index value is %d and the value is %c \n", index, value)
	}

}