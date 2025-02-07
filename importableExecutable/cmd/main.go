package main

import (
	"fmt"

	Calculator "github.com/ArchitJain2134/gocode/importableExecutable/calculator"
	// "github.com/ArchitJain2134/gocode/importableExecutable/store"
	// above line is the syntax to import other package from the directory
)

func main() {
	fmt.Println("sum=", Calculator.Add(5, 8))
	usr:= Calculator.User{}
	usr.Name="archit"

	// var dboprs store.DBOperations
	// dboprs=store.Store{}
	// dboprs.SaveRecord("just print this record....")

}
