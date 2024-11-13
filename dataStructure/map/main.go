/*
In GO, map is a data structure that provide an unordered collection of key-value pairs
where each key must be unique.
It is similar to dictionaries or hash map in other programming lang
Maps are used to associate values with keys and allow for efficient retrieval of values based on these keys
*/
package main

import "fmt"

func main() {
	// string<--> int   ie    name<--> grades
	// here we use name as a key and grade as a value
	studentGrades := make(map[string]int)
	// entering students data
	studentGrades["archit"] = 100
	studentGrades["aryan"] = 35
	studentGrades["mehak"] = 76
	studentGrades["sneha"] = 86
	studentGrades["kartik"] = 83

	// searching the data from the database
	fmt.Println("Marks of aryan is", studentGrades["aryan"])

	// update value
	studentGrades["aryan"] = 95
	fmt.Println("New Marks of aryan is", studentGrades["aryan"])

	// deleting any data
	delete(studentGrades, "aryan")
	fmt.Println("Marks of aryan is", studentGrades["aryan"])

	// checking if the key exists or not
	number, exist := studentGrades["David"]
	fmt.Println("marks of teh student is:", number)
	fmt.Println("student exist:", exist)

	// we can also use interface here same as array

	for key, value := range studentGrades {
		fmt.Printf("for key %v the value is %v\n", key, value)
	}

}