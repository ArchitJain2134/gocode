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

func main() {
	var vishal UserA
	vishal.Name = "Vishal Pandey"
	vishal.Address = "Delhi"
	vishal.Contact = 1234567890
	sonam := UserB{
		Name:    "Sonam Bajwa",
		Address: "Saharanpur",
		Contact: 7310815154,
	}
	vishal.addUser() // hum direct function jaise call nhi kr skte methods ko
	sonam.addUser()

}

func (user *UserA) addUser() {
	// hum yha pr data ko bhi change kr skte h
	// user.Name="Archit Jain"
	fmt.Println("ye user a wala chal rha h", user)

}
func (user UserB) addUser() {
	fmt.Println("ye user b wala chal rha h", user)

}
func (user UserA) addUserdub() {
	fmt.Println("ye user a ka dub wala chal rha h", user)

}
