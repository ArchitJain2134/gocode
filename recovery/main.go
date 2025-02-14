// It is used to generate error is criteria is not fulfiled or wrong info is given as an input
package main

import (
	"fmt"
	"os"
)

func main() {

	dbconnection := true
	if !dbconnection {
		fatalerror("cannot connect with DB so closing the program")
	}

	// pass := "12346789"
	// if len(pass)<8 {
	// 	error := fmt.Errorf("password does not match our criteria")
	// 	fmt.Println(error)

	// }else{
	// 	savetodb(pass)
	// }

	// normal
	a := []string{
		"a wala save hua hai",
	}
	savetodb(a)

	// panic error
	b := []string{}
	savetodb(b)

	// normal
	c := []string{
		"c wala save hua hai",
	}
	savetodb(c)
}

func recoverpanic() {
	r := recover()
	if r != nil {
		fmt.Printf("kuch to gadbad hui message=%v\n", r)

	}

}

// fatal error

func fatalerror(message string) {
	fmt.Printf("Some fetal error occured message = %v\n", message)
	os.Exit(1)

}

// func panicerror(message string){
// 	fmt.Printf("Some fetal error occured message = %v\n",message)
// 	os.Exit(1)

// }

func savetodb(data []string) {
	defer recoverpanic()
	if len(data) < 1 {
		panic("no record found to save")
	} else {
		fmt.Printf("record=%v", data)
	}

}
