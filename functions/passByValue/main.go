package main

import "fmt"

func main() {
	fmt.Println("Enter your marks in maths")
	var x int
	fmt.Scanln(&x)
	fmt.Println("Enter your marks in science")
	var y int
	fmt.Scanln(&y)
	add := add(x, y)
	fmt.Println("your answer after addition is:", add)
	subtract := subtract(x, y)
	fmt.Println("your answer after subtraction is:", subtract)
	multiply := multiplication(x, y)
	fmt.Println("your answer after subtraction is:", multiply)
	div, err := divide(x, y)
	if err == nil {
		fmt.Printf("your answer is %v", div)
	} else {
		fmt.Println("Your denominator is zero hence divide is not possible")

	}

}
func add(x, y int) int {

	return x + y
}

func subtract(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}

}
func multiplication(x, y int) int {

	return x * y
}
func divide(x, y int) (int, error) {
	if y != 0 {
		return x / y, nil
	} else {
		return 0, fmt.Errorf("dividing with 0 is not possible")
		//agr 2 value return hoti h to aise likh sakte hai
	}

}
