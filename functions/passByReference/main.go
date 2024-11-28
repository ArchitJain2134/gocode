package main

import "fmt"

func main() {
	f := 5
	g := 5
	add(&f, g)

	fmt.Println("f=", f, "g=", g)

}
func add(a *int, b int) {
	*a = *a + 1     // ye wala pass by reference h
	b = b + 1      //  ye wala pass by value h
}
