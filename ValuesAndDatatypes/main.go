package main

import "fmt"

func main() {
	fmt.Println("This is data type program")

	num := 32
	addofnum := &num

	num2 := 5445.344
	addofnum2 := &num2

	flags := true
	addofboolean := &flags

	name := "Archit jain"
	addofstring := &name

	fmt.Printf(" integer number=%v\n float number=%v\n boolean=%v\n string=%v", num, num2, flags, name)
	fmt.Println("addresses")
	fmt.Printf("Adress of int number=%v\nAdress of float number=%v\nAdress of boolean=%v\nAdress of string=%v", addofnum, addofnum2, addofboolean, addofstring)

}
