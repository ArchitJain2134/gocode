package main

import "fmt"

// non primitive data type

type Student struct{
	Name string
	Class int
	Rollnumber int
	Studentaddress Address

}

type Address struct{
	lane1 string
	lane2 string
	post string
	district string
	village string
}
func main() {

	// primitive approach
	// var name_archit string
	// var class_archit int
	// var rollnumber_archit int

	// var name_radhika string
	// var class_radhika int
	// var rollnumber_radhika int

	// -------------------------------------------------------------------------------
	// Non Primitive approach
	var vishal Student
	// var aryan Student
	vishal.Class=12
	vishal.Name="Vishal Singh"
	vishal.Studentaddress.village="rampur"

	// another way of declaration
	shivam:=Student{
		Name: "shivam",
		Class: 11,
		Rollnumber: 112,
		Studentaddress: Address{        //This is called nested structure
			district: "Saharanpur",
			lane1: "NH21",

		},
	}
	var vishaladdress *Student
	vishaladdress=&vishal

	fmt.Println("addresss of vishal is = ",*vishaladdress)
	// * (pointer) symbol ke andr value store hoti or. & symbol me address ki value store hoti h (memory me jagah)

	// INTERFACE: interface is very useful we dont have to defien the type in interfacen it takes the value and understand on its own

	val:=12
	val2:="121212"

	var interfaceExample interface{}
	// interface ka use tab kia jaa skata h jb pta na hoi aamne se konse type ka variable aayega

	interfaceExample=val
	fmt.Println("Interface value=",interfaceExample)

	interfaceExample=val2
	fmt.Println("Interface value=",interfaceExample)
	
	interfaceExample=true
	fmt.Println("Interface value=",interfaceExample)

	fmt.Println("Running my go code.....")
	fmt.Println("hello",vishal)
	fmt.Println("hello",shivam)
}