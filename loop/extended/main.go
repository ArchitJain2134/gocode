package main

import "fmt"

func main() {
	fmt.Println("Hello welcome to my practice session")
	Days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// for i:=0;i<len(Days);i++{
	// 		fmt.Println(Days[i])
	// 	}
	// ---------------------------------------------------------------------------------------------------------
	// for x := range Days {
	// 	// here x gives the index value
	// 	fmt.Println(Days[x])
	// }
	// ---------------------------------------------------------------------------------------------------------
	for index, day := range Days {
		fmt.Printf("Index value is %v and Day is %v\n", index, day)
	}
	// ---------------------------------------------------------------------------------------------------------
	checkvalue := 0
	for checkvalue < 10 {

		if checkvalue==2{
			goto lco
		}

		if checkvalue==5{
			// break 
			//break keyword will stop the program un the middle
			checkvalue++
			continue
			// continue keyword is used to skip the element and continue from the next one
		}
		fmt.Println("Value is: ", checkvalue	)
		checkvalue++
	}

	lco:{
		// here lco: is a label || we can declare any label using :
		fmt.Println("Jumping to www.google.com")
	}
	


}