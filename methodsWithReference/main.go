package main

import "fmt"

// kisi bhi function ke pehle jb hum structure add kr dete h to wo method bn jaata h

type UserA struct {
	Name    string
	Address string
	Contact int
}

type UserB struct {
	Name    string
	Address string
	Contact int
}

type UserOperation interface {
	// interface me hum method declare kr skte h
	addUser()
}

func main() {
	var vishal UserA
	vishal.Name = "Vishal Pandey"
	vishal.Address = "Delhi"
	vishal.Contact = 32547568767
	sonam := UserB{
		Name:    "Sonam ",
		Address: "Saharanpur",
		Contact: 43654765867,
	}

	// vishal.addUser() // hum direct function jaise call nhi kr skte methods ko
	// sonam.addUser()

	var useroperation UserOperation
	useroperation = vishal
	useroperation.addUser()  // ye wala user a ko call kr rha h

	useroperation=sonam
	useroperation.addUser()  // ye wala b wale ko call kr rha h

}

func (user UserA) addUser() {
	// hum yha pr data ko bhi change kr skte h
	// user.Name="Archit Jain"
	fmt.Println("ye user a wala chal rha h", user)

}
func (user UserB) addUser() {
	fmt.Println("ye user b wala chal rha h", user)

}
