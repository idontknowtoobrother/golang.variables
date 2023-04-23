package main

import "fmt"

// Global variable declaration
var global_count int = 0

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
	number_imp_var := 0
	string_imp_var := "this is Implicit variable"
	bool2_var := true

	fmt.Println("number_imp_var: ", number_imp_var)
	fmt.Println("string_imp_var: ", string_imp_var)
	fmt.Println("bool2_var: ", bool2_var)

	// Global variable print
	fmt.Println("global_count: ", global_count)
	global_count++
	// self = self + 1
	fmt.Println("global_count: ", global_count)

	// Call function increase global count
	increaseGlobalCount()
	// Call function print global count
	printGlobalCount()
}

func increaseGlobalCount() {
	global_count++
}

func printGlobalCount() {
	fmt.Println("global_count: ", global_count)
}
