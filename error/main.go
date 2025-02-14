// It is used to generate error is criteria is not fulfiled or wrong info is given as an input
package main

import (
	"fmt"
	"os"
)

func main() {
	pass := "12346789"
	if len(pass)<8 {
		error := fmt.Errorf("password does not match our criteria")
		fmt.Println(error)
		
	}else{
		savetodb(pass)
	}
	dbconnection:=true
		if !dbconnection{
			fatalerror("cannot connect with DB so closing the program")
		}
	}

	// fatal error

	func fatalerror(message string){
		fmt.Printf("Some fetal error occured message = %v",message)
		os.Exit(1)

	}


	func savetodb (data string){
		fmt.Printf("sava to data base record = %v",data)
	
		
	}
	// there are 2 types of error
	// 1: fatal error (it stops execution )
	// 2: panic error
