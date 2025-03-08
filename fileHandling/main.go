// os package iis used for file handling in golang

package main

import (
	"fmt"
	"os"
)

func main() {
	// creating a new file

	file, err := os.Create("newfile.txt")

	if err != nil {
		fmt.Println("Error while creating a file")
		return
	}

	fmt.Println("File created successfully")
	defer file.Close()

	// writting text to a file
	written, err := file.WriteString("hello,my name is archit jain")
	if err != nil {
		fmt.Println("Error while writing to a file")
		return
	}
	fmt.Printf("written %v to a file", written)
	readfile("newfile.txt")

	// appending to a file
	append("newfile.txt")
	readfile("newfile.txt")
	


	// //  deleting a file
	// err:=os.Remove("newfile.txt")
	// if err!= nil{
	// 	fmt.Printf("error while deleting a file %v",err)
	// 	return
	// }
	// fmt.Println("File deletion successfully")


	
}

func readfile(filename string) {

	data,err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while writing to a file")
		return
	}
	fmt.Printf("\nData in the file:  %v \n",string(data))
}

func append (filename string){
	file,err:=os.OpenFile(filename, os.O_APPEND|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Println("error while appending a file")
		return
	}
	defer file.Close()

	_,error:=file.WriteString("\n This is the appended line")
	if error!=nil{
		fmt.Println("error while appending a file part 2")
		return
	}
	fmt.Println("\n appending data successfully")
	defer file.Close()
}

// for removing a file

// func remove (filename string){
// 	err:=os.Remove(filename)
// 	if err!=nil{
// 		fmt.Printf("\n error while deleting a file %v", err)
// 		return
// 	}
// 	fmt.Println("File removed successfully")
// }
