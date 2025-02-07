package Calculator

// ye wala code khus nhi chalega main code me iss part/ fiel ko import kr skte h

type User struct{
	Name string
	contactNo int
	// name is importable but concatcNo is not
}

func Add(a int, b int) int {
	// function ka naam capital letter se start hoga tbhi doosre file me use kr skte h wrna expoer nhi ho skta
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}
