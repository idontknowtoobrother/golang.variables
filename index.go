package main

import "fmt"

func main() {
	fmt.Println("all of variables in golang")

	// Explicitly declare variable type
	// number_var type int and value 0
	var number_var int = 0
	// string_var type string and value "this is string"
	var string_var string = "this is string"
	// bool_var type bool and value false
	var bool_var bool = false

	fmt.Println("number_var: ", number_var)
	fmt.Println("string_var: ", string_var)
	fmt.Println("bool_var: ", bool_var)

	// Implicitly declare variable type
	bool2_var := true
	fmt.Println("bool2_var: ", bool2_var)
}
