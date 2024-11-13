package main

import "fmt"

func main() {
	var list [5]int // iss line ka mtlb h list name ka ek array h jisme 5(0-4) integer type ki values store ho skti h

	list[0] = 10
	list[1] = 11
	list[2] = 12
	list[3] = 13
	list[4] = 14

	fmt.Println("list = ", list)

	// we can also over write the element in array

	list[3]=43
	fmt.Println("list = ", list)

	// similarly we can also create an array for string type

	var name [5]string // iss line ka mtlb h list name ka ek array h jisme 5(0-4) integer type ki values store ho skti h

	name[0] = "Archit jain"
	name[1] = "Aryan pandey"
	name[2] = "Saanvi singhal"
	name[3] = "Mehak"
	name[4] = "Unnati Tiwari"

	fmt.Println("names = ", name)

	// in string also we can over write an element of an array

	name[2]="Arpita yadav"
	fmt.Println("names = ", name)

	// if we want to store multiple data types in a single array then we use array of interface

	var unique [5] interface{}  //syntax {} curly braces lgane jaruri hai

	unique[0]=12
	unique[1]="jai ho"
	unique[2]=42
	unique[3]="thank you"
	unique[4]="Halla bol"
	fmt.Println("unique",unique)

	// SLICES!!!!!!!!!!!!!
	// slice ka ek main use h ki usme hum define nhi krte ki kitni value store karega jitni hum value rakhna chaahe utni rakh skte hai
	var students [] int

	// students = append(students, ) ye slices ke andr value store krna ka syntax h bracket ke andr jo students likha h uska mtlb h pehle se jo value h unhe same to same copy karo or comma ke aage jo nayi values daalni h wo likhenge
	 students = append(students, 23, 34,43,56)
	 fmt.Println("TOtal students= ",students)
}
