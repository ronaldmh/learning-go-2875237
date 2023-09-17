package main

import (
	"fmt"
)

const aConst int = 64

func main() {
	var aString = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variable type is %T\n", aString)
	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another String\n"

	fmt.Println(anotherString)
	fmt.Printf("The variable type is %T\n", anotherString)

	var anotherInt = 53

	fmt.Println(anotherInt)
	fmt.Printf("The variable type is %T\n", anotherInt)

	fmt.Println(aConst)
	fmt.Printf("The variable type is %T\n", aConst)

}
