package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	a = 332

	var b int16
	b = int16(a)      //this helps in converting int type to int16 type
	fmt.Println("b= ",b)

	var c string
	c="123123"
	num,err:=strconv.Atoi(c)   //This helps in converting string type to int type if possible
	fmt.Println("number=",num,"\nerror=",err)

	age:= 43
	d:=fmt.Sprintf("age=%v",age)  // Sprintf function hels in changing any format to string type
	fmt.Println("age =%d",d)



}