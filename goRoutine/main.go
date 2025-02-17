package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 11; i++ {
		fmt.Printf("a wala func chal rha h , i=%v\n",i)
	}
}
func b() {
	for i := 0; i < 11; i++ {
		fmt.Printf("b wala func chal rha h , i=%v\n",i)
	}
}
func c() {
	for i := 0; i < 11; i++ {
		fmt.Printf("c wala func chal rha h , i=%v\n",i)
	}
}
func d() {
	for i := 0; i < 11; i++ {
		fmt.Printf("d wala func chal rha h , i=%v\n",i)
	}
}

func main() {

	// a()
	// b()
	// c()
	// d()
	// go is a keyword used for implementing go routine
	// any function can run any time there is no serial order maintaine sin go routine
	go a()
	go b()
	go c()
	go d()
	time.Sleep(time.Second*3)
	// above line of code is used to wait for 3 seconds before shutting the compiler or before exitting the main func
}