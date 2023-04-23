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

	var arraysOfInt = []int{1, 2, 3, 4, 5}
	var arraysOfString = []string{"one", "two", "three", "four", "five"}
	var arraysOfFloat = []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println("arraysOfInt: ", arraysOfInt)
	fmt.Println("arraysOfString: ", arraysOfString)
	fmt.Println("arraysOfFloat: ", arraysOfFloat)

	var emptyArrayOfFloat [5]float64

	for index, number := range arraysOfInt {
		emptyArrayOfFloat[index] = float64(number) + arraysOfFloat[index]
	}
	fmt.Println("emptyArrayOfFloat: ", emptyArrayOfFloat)

}

func increaseGlobalCount() {
	global_count++
}

func printGlobalCount() {
	fmt.Println("global_count: ", global_count)
}
