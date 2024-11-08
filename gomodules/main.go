package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("this is my modules")
	_ = mux.Route{}
}
