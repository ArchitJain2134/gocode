package main

import "fmt"

func main() {
	fmt.Println("This is my switch program")

	daysofweek := 2

	switch daysofweek {
	case 1:
		{
			fmt.Println("MONDAY")
		}
	case 2:
		{
			fmt.Println("TUESDAY")
		}
	case 3:
		{
			fmt.Println("WEDNESDAY")
		}
	case 4:
		{
			fmt.Println("THURSDAY")
		}
	case 5:
		{
			fmt.Println("FRIDAY")
		}
	case 6:
		{
			fmt.Println("SATURDAY")
		}
	case 7:
		{
			fmt.Println("SUNDAY")
		}
	default:
		{
			fmt.Println("ERROR!!!!!")
		}

	}
}
