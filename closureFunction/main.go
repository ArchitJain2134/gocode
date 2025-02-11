package main
// closure function always works with anonymous function

import "fmt"

// func add() func() int { 
// 	// this makes this the anonymous function
// 	i:=0
// 	return func() int{
// 		// this is the example of nested version
// 		i++
// 		return i

//     }
// }

// func main(){
// 	result:=add()
// 	// this will reutrn the address of the add function which we cal also use later and the previous value remains stored  
// 	fmt.Println("result=", result())
// }



// One of teh use case is thatm suppose 
// we need to give unique id to every customer so we do not need to learn or remember the previous id we can simply write the new line of code and it will give the new uniqe id
// for example
func ID() func() int { 
	// this makes this the anonymous function
	i:=0
	return func() int{
		// this is the example of nested version
		i++
		return i

    }
}

func main(){
	result:=ID()
	// this will reutrn the address of the add function which we cal also use later and the previous value remains stored  
	fmt.Println("Ram id=", result())
	fmt.Println("Archit id=", result())
	fmt.Println("Ayush id=", result())
}

